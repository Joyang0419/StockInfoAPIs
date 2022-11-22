package implements

import (
	"StockInfoAPIs/internals/controllers/dailyController"
)

var (
	DailyController dailyController.InterfaceDailyController = dailyController.NewDailyController(DailyService)
)
