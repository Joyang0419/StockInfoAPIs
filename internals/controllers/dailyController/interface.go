package dailyController

import "StockInfoAPIs/internals/repos/dailyRepo"

type InterfaceDailyController interface {
	GetDailyPrices(stockCode string, startDate string, endDate string) []dailyRepo.DailyPrice
}
