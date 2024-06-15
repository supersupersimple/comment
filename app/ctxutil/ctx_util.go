package ctxutil

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/supersupersimple/comment/ent"
)

func DB(c *gin.Context) *ent.Client {
	return c.MustGet("db").(*ent.Client)
}

func WithDB(c *gin.Context, client *ent.Client) {
	c.Set("db", client)
}

func WithLoginStatus(c *gin.Context, loggedIn bool) {
	c.Set("loggedIn", loggedIn)
}

func LoggedIn(c context.Context) bool {
	if ginCtx, ok := c.(*gin.Context); ok {
		return ginCtx.MustGet("loggedIn").(bool)
	}
	if loggedIn, ok := c.Value("loggedIn").(bool); ok {
		return loggedIn
	}
	return false
}
