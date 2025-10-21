package main

import (
	"B00k/api"
	"B00k/service"

	"github.com/gin-gonic/gin"
)

func main() {
	service.Service.Init()

	r := gin.Default()

	r.GET("/book/:id", api.GetBookInfo())
	r.GET("/user", api.GetUserInfo())
	r.GET("/book/list", api.GetMultipleBookInfo())
	r.GET("/book/count", api.GetSize())
	r.POST("/book/add", api.AddBook())
	r.POST("/user/register", api.RegisterNewUser())
	r.DELETE("/book/remove/:id", api.DeleteBook())
	r.DELETE("/user/remove/:id", api.DeleteUser())
	r.PATCH("/book/update", api.UpdateBook())
	r.PATCH("/user/update", api.UpdateUser())

	r.Run()
}
