package model

import (
	"strconv"
)

type BookInfo struct {
	Id       int    `json:"id" gorm:"primaryKey;type:int:column:id"`
	Title    string `json:"title" gorm:"not null;type:varchar(256);column:title"`
	Author   string `json:"author" gorm:"not null;type:varchar(256);column:author"`
	Rating   int    `json:"rating" gorm:"not null;type:int;column:rating"`
	CoverUrl string `json:"url" gorm:"-"`
}

func (binfo *BookInfo) ToString() string {
	return "{Book: \"" + strconv.Itoa(int(binfo.Id)) + ": " + binfo.Title + "\", " +
		"Author: " + binfo.Author + ", " +
		"Rating: " + strconv.Itoa(binfo.Rating) + "}"
}
