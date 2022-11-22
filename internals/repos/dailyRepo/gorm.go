package dailyRepo

import (
	"StockInfoAPIs"
	"StockInfoAPIs/internals/helpers"
	"StockInfoAPIs/internals/models"
	"StockInfoAPIs/internals/tools/dbManager"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type GormDailyRepo struct {
	DBManager dbManager.InterfaceDBManger
}

func NewGormDailyRepo(DBManager dbManager.InterfaceDBManger) *GormDailyRepo {
	return &GormDailyRepo{
		DBManager: DBManager,
	}
}

func (repo *GormDailyRepo) GetCalculateTimestamp(table string) time.Time {
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	var calculateTimestamp int64
	db.Model(models.CalculateTimestampModel{Table: table}).Pluck("calculate_timestamp", &calculateTimestamp)

	if calculateTimestamp == 0 {
		return time.Time{}
	}
	return time.Unix(calculateTimestamp, 0)
}

func (repo *GormDailyRepo) GetDailyPrices(stockCode string, startDate string, endDate string) (
	result []DailyPrice) {
	if status := helpers.FitDateFormat(StockInfoAPIs.DateStringFormat, startDate, endDate); !status {
		log.Panic("correct dateFormat: ", StockInfoAPIs.DateStringFormat)
	}
	db := repo.DBManager.ProvideDBConnection().(*gorm.DB)
	query := db.Table("basic").Select("basic.stock_code, basic.stock_name, category.name as category_name").
		Joins("left join category on basic.category_id = category.id")

	mapConditions := make(map[string]interface{})
	mapConditions["daily_price.stock_code"] = stockCode

	db.Table("daily_price").Select(
		"daily_price.stock_code, daily_price.volume, "+
			"daily_price.opening_price, daily_price.closing_price, daily_price.highest_price, "+
			"daily_price.lowest_price, daily_price.calculate_date, q.stock_name, q.category_name").Joins(
		"left join (?) q on daily_price.stock_code = q.stock_code", query).Where(mapConditions).Where(
		"daily_price.calculate_date BETWEEN ? AND ?",
		startDate, endDate).Scan(&result)
	return result
}
