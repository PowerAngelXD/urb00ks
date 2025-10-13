package model

import (
	"B00k/middleware"
	"fmt"
)

type B0Lib struct {
	LibName    string
	LatestDate string
	Books      map[string]BookInfo
	Recommends map[string]BookInfo
	Users      map[string]UserInfo
}

func (lib *B0Lib) PrintAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book.ToString())
	}
}

func (lib *B0Lib) Init() {
	lib.Books = make(map[string]BookInfo)
	lib.Users = make(map[string]UserInfo)
	lib.Recommends = make(map[string]BookInfo)

	middleware.LibraryLog("Initialize the library's maps... Done")
}

var Lib = B0Lib{LibName: "Official Lib", LatestDate: "2025-10-12"}
