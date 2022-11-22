package implements

import (
	"StockInfoAPIs/internals/services/dailyService"
)

var (
	DailyService dailyService.InterfaceDailyService = dailyService.NewDailyService(DailyRepo)
)
