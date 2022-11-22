package dailyController

import (
	"StockInfoAPIs/internals/repos/dailyRepo"
	"StockInfoAPIs/internals/services/dailyService"
)

type DailyController struct {
	DailyService dailyService.InterfaceDailyService
}

func NewDailyController(dailyService dailyService.InterfaceDailyService) *DailyController {
	return &DailyController{
		DailyService: dailyService,
	}
}

func (controller *DailyController) GetDailyPrices(stockCode string, startDate string, endDate string) (result []dailyRepo.DailyPrice) {
	return controller.DailyService.GetDailyPrices(stockCode, startDate, endDate)
}
