package dailyRouter

import (
	"StockInfoAPIs/internals/repos/dailyRepo"
	"StockInfoAPIs/internals/routers"
)

type DailySuccessResponse struct {
	StatusCode int `example:"200"`
	Data       []dailyRepo.DailyPrice
	routers.BaseRouteResponse
}

type DailyErrorResponse struct {
	StatusCode int    `example:"400"`
	Data       string `example:"Required params: [stock_code, start_date, end_date]"`
	routers.BaseRouteResponse
}
