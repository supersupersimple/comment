package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/supersupersimple/comment/app/config"
)

func FirstSetup() gin.HandlerFunc {
	return func(c *gin.Context) {
		if viper.GetBool(config.Setup) {
			c.Next()
		} else {
			c.Redirect(http.StatusFound, "/admin/setup")
		}
	}
}
