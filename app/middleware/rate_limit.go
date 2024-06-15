package middleware

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

func LeakBucket(limiter ratelimit.Limiter) gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limiter.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}
