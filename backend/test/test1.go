package test

import (
	"B00k/middleware"
	"B00k/model"
	"B00k/service"
)

func Test1() {
	a := model.BookInfo{Id: 1234, Title: "hello", Author: "Me", Rating: 7, CoverUrl: "http"}
	service.Service.LibSv.AddBookByInstance(a)

	if service.Service.LibSv.IsBookExist("hello") {
		book, _ := service.Service.LibSv.GetBook("hello")
		middleware.TestLog("book 'hello' is existed, Author: " + book.Author)
	} else {
		middleware.TestLog("book 'hello' is not existed")
	}
}
