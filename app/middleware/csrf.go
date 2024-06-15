package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/justinas/nosurf"
	"github.com/supersupersimple/comment/app/model"
	"github.com/supersupersimple/comment/app/templates/comment/components"
)

func CSRF() gin.HandlerFunc {
	nextHandler, wrapper := adapter.New()
	ns := nosurf.New(nextHandler)
	ns.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "" {
			components.ReplyBox(nil, model.ReplyStatusInvalidCSRF).Render(r.Context(), w)
		} else {
			http.Error(w, "Custom error about invalid CSRF tokens", http.StatusBadRequest)
		}
	}))
	return wrapper(ns)
}
