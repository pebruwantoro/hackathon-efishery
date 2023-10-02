package driver

import (
	"fmt"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"

	"github.com/pebruwantoro/hackathon-efishery/config"
)

func NewPostgreSQLDatabase(cfg config.DBConfig) (db *gorm.DB, err error) {
	fmt.Println("Try NewDatabase ...")

	dsn := cfg.GetDSN()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	if cfg.DebugMode {
		db = db.Debug()
	}

	return
}
