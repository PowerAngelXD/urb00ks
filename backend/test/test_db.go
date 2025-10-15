package test

import (
	"B00k/dao"
	"B00k/logger"
	"B00k/model"
)

func TestDB() {
	book := model.BookInfo{Title: "hello", Author: "Me", Rating: 7, ImageUrl: "http"}
	dao.OfficialDB.Create(&book)
	user := model.UserInfo{Name: "Liuj", Birth: "2020-03-04", Password: "hello-world"}
	dao.OfficialDB.Create(&user)
	logger.TestLog("Create bookinfo done!")
}
