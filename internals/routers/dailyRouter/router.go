package dailyRouter

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	dailyPrice := router.Group("daily_price")
	dailyPrice.GET("", GetDailyPriceHandler)
}
