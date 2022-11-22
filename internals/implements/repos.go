package implements

import (
	"StockInfoAPIs/internals/repos/dailyRepo"
)

var (
	DailyRepo dailyRepo.InterfaceDailyRepo = dailyRepo.NewGormDailyRepo(GormDBManager)
)
