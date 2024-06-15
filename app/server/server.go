package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/supersupersimple/comment/app/api"
	"github.com/supersupersimple/comment/app/config"
	"github.com/supersupersimple/comment/app/middleware"
	"github.com/supersupersimple/comment/app/render"
	"github.com/supersupersimple/comment/ent"
	"go.uber.org/ratelimit"
	_ "modernc.org/sqlite"
)

func StartWebServer() {
	client := initDB()
	defer client.Close()

	config.LoadConfig(client)

	r := gin.Default()
	r.Use(middleware.CSRF())
	r.Use(middleware.Ent(client))

	r.GET("/admin/setup", api.AdminSetup)
	r.POST("/admin/setup", api.AdminSetup)

	limiter := ratelimit.New(1)
	r.Use(cors.New(corsConfig()))

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &render.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	r.Static("/dist", "./app/assets/dist")
	r.Use(middleware.FirstSetup())
	r.Use(sessions.Sessions("sssc", cookie.NewStore([]byte(viper.GetString(config.Session)))))

	public := r.Use(middleware.Auth(false))
	public.GET("/comments", api.GetComments)
	public.POST("/comment", api.AddComment, middleware.LeakBucket(limiter))

	public.GET("/admin/login", api.AdminLogin)
	public.POST("/admin/login", api.AdminLogin)
	private := r.Use(middleware.Auth(true))
	private.GET("/admin/comments", api.AdminGetComments)
	private.POST("/admin/approve/:id", api.ApproveComment)
	private.POST("/admin/reject/:id", api.RejectComment)

	r.Run(":8080")
}

func initDB() *ent.Client {
	db, err := sql.Open("sqlite", "file:comments.sqlite?cache=shared&_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	db.SetMaxOpenConns(1)
	drv := entsql.OpenDB("sqlite3", db)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func corsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = viper.GetStringSlice(config.AllowOrigins)
	fmt.Println(corsConfig.AllowOrigins)
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	return corsConfig
}
