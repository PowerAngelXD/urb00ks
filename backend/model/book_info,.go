package model

import (
	"strconv"
)

type BookInfo struct {
	Id       int64  `json:"id" gorm:"primaryKey;type:int:column:id;autoIncrement"`
	Title    string `json:"title" gorm:"not null;type:varchar(256);column:title"`
	Author   string `json:"author" gorm:"not null;type:varchar(256);column:author"`
	Rating   int    `json:"rating" gorm:"not null;type:int;column:rating"`
	Views    int    `json:"views" gorm:"type:int;column:views"`
	ImageUrl string `json:"url" gorm:"type:text;column:image_url"`
}

func (BookInfo) TableName() string {
	return "book_list"
}

func (binfo *BookInfo) ToString() string {
	return "{Book: \"" + strconv.FormatInt(binfo.Id, 10) + ": " + binfo.Title + "\", " +
		"Author: " + binfo.Author + ", " + "Views: " + strconv.Itoa(binfo.Views) + ", " +
		"Rating: " + strconv.Itoa(binfo.Rating) + "}"
}
