package dao

import (
	"B00k/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var OfficialDB *gorm.DB

func init() {
	logger.DBLog("Start connecting the database...")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		logger.DBLog("Connect the database failed, details: " + err.Error())
	}

	DB, err := db.DB()
	if err != nil {
		logger.DBLog("Get the basic database failed, details: " + err.Error())
	}

	logger.DBLog("Initializing the IdleTime And MaxOpen settings...")
	DB.SetConnMaxIdleTime(time.Hour)
	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(100)

	OfficialDB = db
	logger.DBLog("Everything was Done!")
}
