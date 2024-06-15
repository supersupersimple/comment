package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/supersupersimple/comment/app/ctxutil"
)

func Auth(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil && required {
			// Abort the request with the appropriate error code
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		ctxutil.WithLoginStatus(c, user != nil)
		// Continue down the chain to handler etc
		c.Next()
	}
}
