package main

import (
	"B00k/api"
	"B00k/dao"
	"B00k/logger"
	"B00k/middleware"
	"B00k/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var isDev = true

func main() {
	dao.RdbInit()
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

	publicGroup := r.Group("/public")

	publicGroup.POST("/register", api.RegisterNewUser())
	publicGroup.GET("/login", api.UserLogin())

	apiGroup := r.Group("/api")
	userGroup := apiGroup.Group("/user")
	bookGroup := apiGroup.Group("/book")

	userGroup.Use(middleware.JWTAuthHandler())

	userGroup.GET("/fetch", api.GetUserInfo())
	userGroup.DELETE("/remove/:id", api.DeleteUser())
	userGroup.PATCH("/update", api.UpdateUser())

	bookGroup.GET("/:id", api.GetBookInfo())
	bookGroup.GET("/list", api.GetMultipleBookInfo())
	bookGroup.GET("/search/:target", api.SearchBooks())
	bookGroup.GET("/count", api.GetSize())
	bookGroup.POST("/add", api.AddBook())
	bookGroup.DELETE("/remove/:id", api.DeleteBook())
	bookGroup.PATCH("/update", api.UpdateBook())

	r.Run(":8080")
}
