package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/supersupersimple/comment/app/api"
	"github.com/supersupersimple/comment/app/config"
	"github.com/supersupersimple/comment/app/middleware"
	"github.com/supersupersimple/comment/app/render"
	"github.com/supersupersimple/comment/ent"
	"github.com/supersupersimple/litestream-lib/lslib"
	"go.uber.org/ratelimit"

	_ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

func StartWebServer(https bool, host string, port int) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	client, lsdb := initDB(ctx)
	defer closeDB(ctx, client, lsdb)

	config.LoadConfig(client)

	r := gin.Default()
	r.Static("/dist", "./app/assets/dist")

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &render.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	limiter := ratelimit.New(viper.GetInt(config.RateLimit))

	r.Use(middleware.CSRF())
	r.Use(middleware.Ent(client))

	r.GET("/admin/setup", api.AdminSetup)
	r.POST("/admin/setup", api.AdminSetup)

	r.Use(cors.New(corsConfig()))

	r.Use(middleware.FirstSetup())
	r.Use(sessions.Sessions("sssc", cookie.NewStore([]byte(viper.GetString(config.Session)))))

	public := r.Use(middleware.Auth(false))
	public.GET("/comments", api.GetComments)
	public.POST("/comment", api.AddComment, middleware.LeakBucket(limiter))
	public.GET("/comment/approve/:token", api.AnonymousApprove)
	public.GET("/comment/reject/:token", api.AnonymousReject)

	public.GET("/admin/login", api.AdminLogin)
	public.POST("/admin/login", api.AdminLogin)
	private := r.Use(middleware.Auth(true))
	private.GET("/admin/comments", api.AdminGetComments)
	private.POST("/admin/approve/:id", api.ApproveComment)
	private.POST("/admin/reject/:id", api.RejectComment)

	if https {
		autotls.RunWithContext(ctx, r, host)
	} else {
		runWithCtx(ctx, r, fmt.Sprintf("%s:%d", host, port), stop)
	}
}

func runWithCtx(ctx context.Context, r *gin.Engine, addr string, stop context.CancelFunc) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	slog.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	slog.Info("Server exiting")
}

func initDB(ctx context.Context) (*ent.Client, lslib.DB) {
	lsdb := lslib.NewDB(lslib.NewConfig("data/comments.sqlite").WithDriverName("sqlite"))
	db, err := lsdb.Open(ctx)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	db.SetMaxOpenConns(1)
	_, err = db.Exec("PRAGMA foreign_keys=ON")
	if err != nil {
		log.Fatalf("failed enable foreign keys: %v", err)
	}
	_, err = db.Exec("PRAGMA journal_mode=WAL")
	if err != nil {
		log.Fatalf("failed enable WAL mode: %v", err)
	}

	drv := entsql.OpenDB("sqlite3", db)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, lsdb
}

func closeDB(ctx context.Context, client *ent.Client, lsdb lslib.DB) {
	client.Close()
	lsdb.Close(ctx)
}

func corsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = viper.GetStringSlice(config.AllowOrigins)
	fmt.Println(corsConfig.AllowOrigins)
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	return corsConfig
}
