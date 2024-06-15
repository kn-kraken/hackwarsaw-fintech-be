package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
