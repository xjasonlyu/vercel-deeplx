package api

import (
	"net/http"

	deeplx "github.com/OwO-Network/DeepLX/service"
	"github.com/gin-gonic/gin"
)

var (
	cfg *deeplx.Config
	app *gin.Engine
)

func init() {
	cfg = deeplx.InitConfig()
	app = deeplx.Router(cfg)
}

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

var _ http.HandlerFunc = Entrypoint
