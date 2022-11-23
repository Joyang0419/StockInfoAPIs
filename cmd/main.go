package main

import (
	_ "StockInfoAPIs/inits"
	"StockInfoAPIs/internals"
	"StockInfoAPIs/internals/middleware"
	"StockInfoAPIs/internals/routers/dailyRouter"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @title StockInfoAPIs
// @version 1.0
// @description StockInfoAPIs swagger
// @host 0.0.0.0:8080
// schemes http
func main() {
	middlewares := []gin.HandlerFunc{
		middleware.CorsMiddleware(),
		gin.Recovery(), // Recovery middleware recovers from any panics and writes a 500 if there was one.
		middleware.JSONLogMiddleware(true, log.DebugLevel),
	}

	server := internals.NewServer(
		gin.DebugMode, []func(router *gin.Engine){dailyRouter.Route}, middlewares...)
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
