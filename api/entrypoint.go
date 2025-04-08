package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	app = gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
