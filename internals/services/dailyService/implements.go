package dailyService

import (
	"StockInfoAPIs/internals/repos/dailyRepo"
)

type DailyService struct {
	DailyRepo dailyRepo.InterfaceDailyRepo
}

func NewDailyService(dailyRepo dailyRepo.InterfaceDailyRepo) *DailyService {
	return &DailyService{
		DailyRepo: dailyRepo,
	}
}

func (service *DailyService) GetDailyPrices(stockCode string, startDate string, endDate string) (result []dailyRepo.DailyPrice) {
	repoResponse := service.DailyRepo.GetDailyPrices(stockCode, startDate, endDate)
	return repoResponse
}
