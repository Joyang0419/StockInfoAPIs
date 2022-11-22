package implements

import (
	"StockInfoAPIs"
	"StockInfoAPIs/internals/tools/dbManager"
	"fmt"
	"gorm.io/driver/mysql"
)

var (
	sqlUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
		StockInfoAPIs.Config.Db.UserName,
		StockInfoAPIs.Config.Db.Password,
		StockInfoAPIs.Config.Db.Host,
		StockInfoAPIs.Config.Db.Port,
		StockInfoAPIs.Config.Db.DbName,
	)
	gormSetting = dbManager.NewGORMDBMSetting(mysql.Open(sqlUrl),
		dbManager.DBMaxIdleConns(StockInfoAPIs.Config.Db.MaxIdleConnection),
		dbManager.DBMaxOpenConns(StockInfoAPIs.Config.Db.MaxOpenConnection),
		dbManager.ConnMaxLifeTimeMinutes(StockInfoAPIs.Config.Db.ConnMaxLifetimeMinutes))
	GormDBManager dbManager.InterfaceDBManger = dbManager.NewGormDBManager(gormSetting)
)
