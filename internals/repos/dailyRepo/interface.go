package dailyRepo

import (
	"github.com/shopspring/decimal"
	"time"
)

type DailyPrice struct {
	StockCode     string          `example:"1101"`                      // 股票代碼
	StockName     string          `example:"台泥"`                        // 股票名稱
	CategoryName  string          `example:"水泥工業"`                      // 類別名稱
	Volume        decimal.Decimal `example:"18471"`                     //交易量(張)
	OpeningPrice  decimal.Decimal `example:"43.8"`                      // 開盤價
	ClosingPrice  decimal.Decimal `example:"44.1"`                      // 收盤價
	HighestPrice  decimal.Decimal `example:"44.15"`                     // 最高價
	LowestPrice   decimal.Decimal `example:"43.8"`                      // 最低價
	CalculateDate time.Time       `example:"2020-01-02T00:00:00+08:00"` // 價格日期
}

type InterfaceDailyRepo interface {
	GetDailyPrices(stockCode string, startDate string, endDate string) []DailyPrice
	GetCalculateTimestamp(table string) time.Time
}
