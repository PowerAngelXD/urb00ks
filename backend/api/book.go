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
				c.JSON(http.StatusNotAcceptable, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 0})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[model.BookInfo]{Status: http.StatusOK, Msg: "Get bookinfo successfully", Data: book})
			}
		}
	}
}
