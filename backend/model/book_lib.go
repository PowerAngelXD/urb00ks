package model

import (
	"B00k/middleware"
	"fmt"
)

type B0Lib struct {
	LibName    string
	LatestDate string
	Books      map[int64]BookInfo
	Recommends map[int64]BookInfo
	Users      map[int64]UserInfo
}

func (lib *B0Lib) PrintAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book.ToString())
	}
}

func (lib *B0Lib) Init() {
	lib.Books = make(map[int64]BookInfo)
	lib.Users = make(map[int64]UserInfo)
	lib.Recommends = make(map[int64]BookInfo)

	middleware.LibraryLog("Initialize the library's maps... Done")
}

var Lib = B0Lib{LibName: "Official Lib", LatestDate: "2025-10-12"}
