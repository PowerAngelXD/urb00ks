package api

import (
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Param("title")
		if service.IsUserExist(&model.Lib, title) {
			book, err := service.GetBook(&model.Lib, title)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusNotAcceptable, ErrorStruct)
			} else {
				c.JSON(http.StatusOK, book)
			}
		}
	}
}
