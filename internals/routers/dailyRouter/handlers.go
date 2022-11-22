package dailyRouter

import (
	"StockInfoAPIs/internals/helpers"
	"StockInfoAPIs/internals/implements"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDailyPriceHandler
// @Router /daily_price [get]
// @Accept  json
// @Produce  json
// @param stock_code query string true "股票代碼" default(1101)
// @param start_date query string true "開始日期" default(2020-01-01)
// @param end_date query string true "結束日期" default(2020-01-02)
// @Success 200 {object} DailySuccessResponse
// @Failure 400 {object} DailyErrorResponse
func GetDailyPriceHandler(c *gin.Context) {
	var (
		stockCode string
		startDate string
		endDate   string
		ok        bool
	)
	stockCode = c.Query("stock_code")
	startDate = c.Query("start_date")
	endDate = c.Query("end_date")

	ok = helpers.IsParamsExist(stockCode, startDate, endDate)

	if !ok {
		errResponse := DailyErrorResponse{StatusCode: http.StatusBadRequest,
			Data: "Required params: [stock_code, start_date, end_date]"}
		errResponse.GinErrorResponse(c, errResponse)
		return
	}

	result := implements.DailyController.GetDailyPrices(stockCode, startDate, endDate)

	successResponse := DailySuccessResponse{StatusCode: http.StatusOK, Data: result}
	successResponse.GinSuccessResponse(c, successResponse)
}
