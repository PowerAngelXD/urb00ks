package test

import (
	"B00k/middleware"
	"B00k/service"
)

func Test1() {
	service.Service.LibSv.AddBook("TestBook", "FZSGBall")

	if service.Service.LibSv.IsBookExistByTitle("hello") {
		book, _ := service.Service.LibSv.GetBook("TestBook")
		middleware.TestLog("book 'hello' is existed, Author: " + book.Author)
	} else {
		middleware.TestLog("book 'hello' is not existed")
	}
}
