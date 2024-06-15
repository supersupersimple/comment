package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/supersupersimple/comment/app/ctxutil"
	"github.com/supersupersimple/comment/ent"
)

func Ent(entClient *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxutil.WithDB(c, entClient)
		c.Next()
	}
}
