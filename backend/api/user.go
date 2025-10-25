package api

import (
	"B00k/logger"
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GET
// 获得userinfo
func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		type_str := c.Query("type")
		content := c.Query("content")
		password := c.Query("password")

		if type_str == "id" {
			logger.ServiceLog("enter id case")
			id, _ := strconv.ParseInt(content, 10, 64)
			logger.ServiceLog("id: " + content)
			if service.Service.UserSv.IsUserExist(id) {
				user, err := service.Service.UserSv.GetUserById(id)
				if err != nil {
					log.Println(err.Error())
					c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 1})
				} else {
					if password != user.Password {
						c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed: password wrong", Data: 1})
					} else {
						c.JSON(http.StatusOK, ReturnStruct[model.UserInfo]{Status: http.StatusOK, Msg: "Get userinfo successfully", Data: user})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Cannot get an unknown user", Data: 1})
			}
		}
		if type_str == "name" {
			if service.Service.UserSv.IsUserExistByName(content) {
				user, err := service.Service.UserSv.GetUser(content)
				if err != nil {
					log.Println(err.Error())
					c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 1})
				} else {
					if password != user.Password {
						c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed: password wrong", Data: 1})
					} else {
						c.JSON(http.StatusOK, ReturnStruct[model.UserInfo]{Status: http.StatusOK, Msg: "Get userinfo successfully", Data: user})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Cannot get an unknown user", Data: 1})
			}
		}
		if type_str == "favs" {
			// 如果是favs type，那么必须用id访问
			id, _ := strconv.ParseInt(content, 10, 64)
			if service.Service.UserSv.IsUserExist(id) {
				user, err := service.Service.UserSv.GetUserById(id)
				if err != nil {
					log.Println(err.Error())
					c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed, details: \"" + err.Error() + "\"", Data: 1})
				} else {
					if password != user.Password {
						c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed: password wrong", Data: 1})
					} else {
						var books []model.BookInfo = make([]model.BookInfo, 0)
						for _, book_id_str := range user.Favs {
							book_id, _ := strconv.ParseInt(book_id_str, 10, 64)
							book, _ := service.Service.LibSv.GetBookById(book_id)
							books = append(books, book)
						}
						c.JSON(http.StatusOK, ReturnStruct[[]model.BookInfo]{Status: http.StatusOK, Msg: "Get userinfo successfully", Data: books})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Cannot get an unknown user", Data: 1})
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

		result := service.Service.UserSv.AddUser(name, birth, pswd) // TODO: 之后考虑加密解密的事情

		if result != nil {
			c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Register user failed, details: " + result.Error(), Data: 1})
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
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Delete user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Delete user successfully", Data: 0})
			}
		} else {
			c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Cannot delete an unknown user", Data: 1})
		}
	}
}

// PATCH
// 更新用户特定信息
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		target_id_str := c.Query("target")
		type_str := c.Query("type")
		content := c.Query("content")
		id, _ := strconv.ParseInt(target_id_str, 10, 64)

		switch type_str {
		case "update_name":
			result := service.Service.UserSv.UpdateName(id, content)
			if result != nil {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully update user", Data: 0})
			}
		case "update_pswd":
			result := service.Service.UserSv.UpdatePassword(id, content)
			if result != nil {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully update user", Data: 0})
			}
		case "add_fav":
			result := service.Service.UserSv.AddFav(id, content)
			if result != nil {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully update user", Data: 0})
			}
		case "delete_fav":
			result := service.Service.UserSv.DeleteFav(id, content)
			if result != nil {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully update user", Data: 0})
			}
		case "update_all":
			// e.g. /user/update?target=1&type=update_all&Xiaoming;123456;book1,book2,book3
			fields := strings.Split(content, ";")
			favs := strings.Split(fields[2], ",")

			result := service.Service.UserSv.UpdateAll(id, fields[0], fields[1], favs)

			if result != nil {
				c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update user failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully update user", Data: 0})
			}
		}
	}
}
