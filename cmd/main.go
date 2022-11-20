package main

import (
	"StockInfoAPIs"
	"StockInfoAPIs/internals/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	middlewares := []gin.HandlerFunc{
		middleware.CorsMiddleware(),
		gin.Recovery(), // Recovery middleware recovers from any panics and writes a 500 if there was one.
		middleware.JSONLogMiddleware(true, log.DebugLevel),
	}
	server := StockInfoAPIs.NewServer(gin.DebugMode, middlewares...)
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
