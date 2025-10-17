package api

import (
	"B00k/model"
	"B00k/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GET
// 获取用户信息
func GetBookInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_str := c.Param("id")
		id, _ := strconv.ParseInt(id_str, 10, 64)
		book, err := service.Service.LibSv.GetBookById(id)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Get userinfo failed, details: " + err.Error(), Data: 1})
		} else {
			c.JSON(http.StatusOK, ReturnStruct[model.BookInfo]{Status: http.StatusOK, Msg: "Get bookinfo successfully", Data: book})
		}
	}
}

// POST
// 增加新的书本
func AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		author := c.Query("author")
		url := c.Query("url")

		result := service.Service.LibSv.AddBook(title, author, url)

		if result != nil {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Add new book failed, details: " + result.Error(), Data: 1})
		} else {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully add a new book", Data: 0})
		}
	}
}

// DELETE
// 删除书本
func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_str := c.Param("id")
		id, _ := strconv.ParseInt(id_str, 10, 64)

		result := service.Service.LibSv.DeleteBook(id)

		if result != nil {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Delete book failed, details: " + result.Error(), Data: 1})
		} else {
			c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusOK, Msg: "Successfully delete a book", Data: 0})
		}
	}
}

// PATCH
// 更新书本信息
func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		target_id_str := c.Query("target")
		type_str := c.Query("type")
		content := c.Query("content")
		id, _ := strconv.ParseInt(target_id_str, 10, 64)

		switch type_str {
		case "update_title":
			result := service.Service.LibSv.UpdateTitle(id, content)
			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update a book failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Successfully update a book", Data: 0})
			}
		case "update_author":
			result := service.Service.LibSv.UpdateAuthor(id, content)
			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update a book failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Successfully update a book", Data: 0})
			}
		case "update_url":
			result := service.Service.LibSv.UpdateUrl(id, content)
			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update a book failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Successfully update a book", Data: 0})
			}
		case "update_rating":
			rt, _ := strconv.ParseInt(content, 10, 32)
			result := service.Service.LibSv.UpdateRating(id, int(rt))
			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update a book failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Successfully update a book", Data: 0})
			}
		case "update_views":
			views, _ := strconv.ParseInt(content, 10, 32)
			result := service.Service.LibSv.UpdateRating(id, int(views))
			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update a book failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Successfully update a book", Data: 0})
			}
		case "update_all":
			// e.g. /user/update?target=1&type=update_all&呼啸山庄;Emily Brontë;10;480;/image/Heathcliff/WildHunt.png
			fields := strings.Split(content, ";")
			rt, _ := strconv.ParseInt(fields[2], 10, 32)
			views, _ := strconv.ParseInt(fields[3], 10, 32)

			result := service.Service.LibSv.UpdateAll(id, fields[0], fields[1], int(rt), int(views), fields[4])

			if result != nil {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Update a book failed, details: " + result.Error(), Data: 1})
			} else {
				c.JSON(http.StatusOK, ReturnStruct[int]{Status: http.StatusBadRequest, Msg: "Successfully update a book", Data: 0})
			}
		}
	}
}
