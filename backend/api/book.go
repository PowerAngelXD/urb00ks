package api

import (
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_str := c.Param("id")
		id, _ := strconv.ParseInt(id_str, 10, 64)
		if service.Service.LibSv.IsBookExist(id) {
			book, err := service.Service.LibSv.GetBookById(id)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusNotAcceptable, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 0})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[model.BookInfo]{Status: http.StatusOK, Msg: "Get bookinfo successfully", Data: book})
			}
		}
	}
}
