package test

import (
	"B00k/model"
	"B00k/service"
	"fmt"
)

func Test1() {
	a := model.BookInfo{Id: 1234, Title: "hello", Author: "Me", Rating: 7, CoverUrl: "http"}
	service.Service.LibSv.AddBookByInstance(a)

	if service.Service.LibSv.IsBookExist("hello") {
		book, _ := service.Service.LibSv.GetBook("hello")
		fmt.Println("[Test] book 'hello' is existed, Author: " + book.Author)
	} else {
		fmt.Println("[Test] book 'hello' is not existed")
	}
}
