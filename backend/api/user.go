package api

import (
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if service.IsBookExist(&model.Lib, name) {
			user, err := service.GetUser(&model.Lib, name)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusNotAcceptable, ErrorStruct)
			} else {
				c.JSON(http.StatusOK, user)
			}
		}
	}
}
