package main

import (
	"B00k/api"
	"B00k/dao"
	"B00k/logger"
	"B00k/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var isDev = true

func main() {
	if isDev {
		logger.DBLog("Start connecting the database...")
		db, err := gorm.Open(mysql.Open(dao.DevDSN), &gorm.Config{})

		if err != nil {
			logger.DBLog("Connect the database failed, details: " + err.Error())
			panic("Connect the database failed, program stop...")
		}

		DB, err := db.DB()
		if err != nil {
			logger.DBLog("Get the basic database failed, details: " + err.Error())
		}

		logger.DBLog("Initializing the IdleTime And MaxOpen settings...")
		DB.SetConnMaxIdleTime(time.Hour)
		DB.SetMaxIdleConns(20)
		DB.SetMaxOpenConns(100)

		dao.OfficialDB = db
		logger.DBLog("Everything was Done!")
	} else {
		if err := dao.ConnectDB(); err != nil {
			os.Exit(1)
		}
	}
	service.Service.Init()

	r := gin.Default()

	r.GET("/api/book/:id", api.GetBookInfo())
	r.GET("/api/user", api.GetUserInfo())
	r.GET("/api/book/list", api.GetMultipleBookInfo())
	r.GET("/api/book/search/:target", api.SearchBooks())
	r.GET("/api/book/count", api.GetSize())
	r.POST("/api/book/add", api.AddBook())
	r.POST("/register", api.RegisterNewUser())
	r.POST("/login", api.UserLogin())
	r.DELETE("/api/book/remove/:id", api.DeleteBook())
	r.DELETE("/api/user/remove/:id", api.DeleteUser())
	r.PATCH("/api/book/update", api.UpdateBook())
	r.PATCH("/api/user/update", api.UpdateUser())

	r.Run()
}
