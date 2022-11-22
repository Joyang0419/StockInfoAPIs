package dailyService

import "StockInfoAPIs/internals/repos/dailyRepo"

type InterfaceDailyService interface {
	GetDailyPrices(stockCode string, startDate string, endDate string) []dailyRepo.DailyPrice
}
