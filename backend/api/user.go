package api

import (
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET
// 获得userinfo
func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_str := c.Param("id")
		id, _ := strconv.ParseInt(id_str, 10, 64)
		if service.Service.UserSv.IsUserExist(id) {
			user, err := service.Service.UserSv.GetUserById(id)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusNotAcceptable, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[model.UserInfo]{Status: http.StatusOK, Msg: "Get userinfo successfully", Data: user})
			}
		} else {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Cannot delete an unknown user", Data: 1})
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

		result := service.Service.UserSv.AddUser(name, birth, pswd) // TODO: 之后考虑加密解密的事情

		if result != nil {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Register user failed, details: " + result.Error(), Data: 1})
		} else {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Register successfully", Data: 0})
		}
	}
}

// DEL
// 删除一个用户
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_str := c.Param("id")
		id, _ := strconv.ParseInt(id_str, 10, 64)
		if service.Service.UserSv.IsUserExist(id) {
			result := service.Service.UserSv.DB.Delete(id)
			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Delete user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Delete user successfully", Data: 0})
			}
		} else {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusNotAcceptable, Msg: "Cannot delete an unknown user", Data: 1})
		}
	}
}

// PATCH
// 更新用户特定信息
