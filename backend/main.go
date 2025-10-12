package main

import (
	"B00k/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	a := model.BookInfo{Id: 1234, Title: "hello", Author: "Me", Rating: 7, CoverUrl: "http"}

	fmt.Println(a.ToString())
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
