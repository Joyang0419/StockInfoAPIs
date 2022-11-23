package internals

import (
	_ "StockInfoAPIs/swaggerDocs" // 匯入swagger docs
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(mode string, routes []func(router *gin.Engine), middlewares ...gin.HandlerFunc) (server *gin.Engine) {
	gin.SetMode(mode)
	server = gin.New()
	server.Use(middlewares...)
	for _, route := range routes {
		route(server)
	}

	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://0.0.0.0:%d/swagger/doc.json", 8080))
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	return server
}
