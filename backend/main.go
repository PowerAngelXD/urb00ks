package main

import (
	"B00k/api"
	"B00k/service"
	"B00k/test"

	"github.com/gin-gonic/gin"
)

func main() {
	test.TestDB()
	service.Service.Init()

	r := gin.Default()

	r.GET("/book/:id", api.GetBookInfo())
	r.GET("/user/:id", api.GetUserInfo())
	r.GET("/book/list", api.GetMultipleBookInfo())
	r.POST("/book/add", api.AddBook())
	r.POST("/user/register", api.RegisterNewUser())
	r.DELETE("/book/remove/:id", api.DeleteBook())
	r.DELETE("/user/remove/:id", api.DeleteUser())
	r.PATCH("/book/update", api.UpdateBook())
	r.PATCH("/user/update", api.UpdateUser())

	r.Run()
}
