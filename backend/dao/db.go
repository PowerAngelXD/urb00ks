package dao

import (
	"B00k/logger"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var OfficialDB *gorm.DB
var DSN string

func ConnectDB() error {
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")

	if DBHost == "" {
		DBHost = "db"
	}
	if DBPort == "" {
		DBPort = "3306"
	}

	DSN = DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	logger.DBLog(fmt.Sprintf("[DataBase] DSN Host is: %s", DBHost))
	logger.DBLog("[DataBase] Start connecting the database...")

	var err error
	const maxRetries = 10
	const delay = 5 * time.Second

	// 4. 重试循环
	for i := 0; i < maxRetries; i++ {
		// 尝试连接
		OfficialDB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

		if err == nil {
			sqlDB, _ := OfficialDB.DB()
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			logger.DBLog("[DataBase] Connect the database successfully.")
			return nil
		}

		time.Sleep(delay)
	}

	if err != nil {
		logger.DBLog("[DataBase] Connect the database failed")
		return err
	}

	sqlDB, _ := OfficialDB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.DBLog("[DataBase] Connect the database successfully.")
	return nil
}
