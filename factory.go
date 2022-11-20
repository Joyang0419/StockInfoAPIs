package StockInfoAPIs

import (
	"github.com/gin-gonic/gin"
)

func NewServer(mode string, middlewares ...gin.HandlerFunc) (server *gin.Engine) {
	gin.SetMode(mode)
	server = gin.New()
	server.Use(middlewares...)
	return server
}
