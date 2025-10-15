package api

import (
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
// 获得userinfo
func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if service.Service.UserSv.IsUserExistByName(name) {
			user, err := service.Service.UserSv.GetUser(name)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusNotAcceptable, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 0})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[model.UserInfo]{Status: http.StatusOK, Msg: "Get userinfo successfully", Data: user})
			}
		}
	}
}

// POST
// 注册一个用户
func RegisterNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		birth := c.Query("birth")
		pswd := c.Query("pswd")

		service.Service.UserSv.AddUser(name, birth, pswd) // TODO: 之后考虑加密解密的事情
		c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Register successfully", Data: 0})
	}
}
