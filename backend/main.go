package main

import (
	"B00k/api"
	"B00k/dao"
	"B00k/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := dao.ConnectDB(); err != nil {
		os.Exit(1)
	}
	service.Service.Init()

	r := gin.Default()

	r.GET("/api/book/:id", api.GetBookInfo())
	r.GET("/api/user", api.GetUserInfo())
	r.GET("/api/book/list", api.GetMultipleBookInfo())
	r.GET("/api/book/count", api.GetSize())
	r.POST("/api/book/add", api.AddBook())
	r.POST("/api/user/register", api.RegisterNewUser())
	r.DELETE("/api/book/remove/:id", api.DeleteBook())
	r.DELETE("/api/user/remove/:id", api.DeleteUser())
	r.PATCH("/api/book/update", api.UpdateBook())
	r.PATCH("/api/user/update", api.UpdateUser())

	r.Run()
}
