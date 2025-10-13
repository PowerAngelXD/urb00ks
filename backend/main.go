package main

import (
	"B00k/test"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	test.TestDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
