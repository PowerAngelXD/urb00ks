package test

import (
	"B00k/logger"
	"B00k/service"
)

func Test1() {
	service.Service.LibSv.AddBook("TestBook", "FZSGBall", "image.png")

	if service.Service.LibSv.IsBookExistByTitle("hello") {
		book, _ := service.Service.LibSv.GetBook("TestBook")
		logger.TestLog("book 'hello' is existed, Author: " + book.Author)
	} else {
		logger.TestLog("book 'hello' is not existed")
	}
}
