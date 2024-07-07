package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"golang.org/x/exp/slog"
)

func LeakBucket(limiter ratelimit.Limiter) gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limiter.Take()
		slog.Info("rate limit", "time delta", now.Sub(prev))
		prev = now
	}
}
