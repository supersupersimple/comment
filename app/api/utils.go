package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/supersupersimple/comment/app/config"
	"github.com/supersupersimple/comment/app/ctxutil"
	"github.com/supersupersimple/comment/app/templates/admin/admincomponents"
	"github.com/supersupersimple/comment/ent"
	"github.com/supersupersimple/comment/ent/comment"
	"golang.org/x/crypto/bcrypt"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ConvertPublishedTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff < time.Second {
		return "just now"
	}

	if diff < 2*time.Second {
		return "1 second ago"
	}

	if diff < 60*time.Second {
		return fmt.Sprintf("%d seconds ago", int(diff/time.Second))
	}

	if diff < 2*time.Minute {
		return "1 minute ago"
	}

	if diff < 60*time.Minute {
		return fmt.Sprintf("%d minutes ago", int(diff/time.Minute))
	}

	if diff < 2*time.Hour {
		return "1 hour ago"
	}

	if diff < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(diff/time.Hour))
	}

	if diff < 2*24*time.Hour {
		return "1 day ago"
	}

	if diff < 7*24*time.Hour {
		return fmt.Sprintf("%d days ago", int(diff/24/time.Hour))
	}

	if diff < 14*24*time.Hour {
		return "1 week ago"
	}

	if diff < 365*24*time.Hour {
		return fmt.Sprintf("%d weeks ago", int(diff/7/24/time.Hour))
	}

	if diff < 2*365*24*time.Hour {
		return "1 year ago"
	}

	return fmt.Sprintf("%d years ago", int(diff/365/24/time.Hour))
}

func SendTgNoti(ctx *gin.Context, co *ent.Comment) {
	tgUrl := viper.GetString(config.TgBotUrl)
	if tgUrl == "" {
		return
	}
	host := viper.GetString(config.Host)
	if host == "" {
		return
	}
	u, err := url.Parse(tgUrl)
	if err != nil {
		return
	}
	q := u.Query()
	chatID := q.Get("chat_id")
	if chatID == "" {
		return
	}

	co, err = ctxutil.DB(ctx).Comment.Query().Where(comment.IDEQ(co.ID)).WithPage().WithUser().Only(ctx)
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	err = admincomponents.TgNotiBox(co.Edges.Page.Title, co.Edges.User.Username, co.Content).Render(ctx, buf)
	if err != nil {
		return
	}
	msgText := buf.String()
	slog.Info("msg", "text", msgText)

	tgUrl = u.Scheme + "://" + u.Host + u.Path

	body := map[string]interface{}{
		"chat_id":    chatID,
		"text":       msgText,
		"parse_mode": "HTML",
		"reply_markup": map[string]interface{}{
			"inline_keyboard": [][]map[string]string{
				{
					{
						"text": "Approve",
						"url":  fmt.Sprintf("%s/comment/approve/%s", host, co.ApproveToken),
					},
					{
						"text": "Reject",
						"url":  fmt.Sprintf("%s/comment/reject/%s", host, co.ApproveToken),
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, tgUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send Telegram notification
	client := &http.Client{}
	slog.Info("send tg noti", "url", req.URL.String())
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("call tg", "err", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("read resp body", "err", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error("call tg failed", "status code", resp.StatusCode, "resp body", string(respBody))
		return
	}
}
