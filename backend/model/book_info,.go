package model

import "strconv"

type BookInfo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Rating   int    `json:"rating"`
	CoverUrl string `json:"url"`
}

func (binfo *BookInfo) ToString() string {
	return "{Book: \"" + strconv.Itoa(binfo.Id) + ": " + binfo.Title + "\", " +
		"Author: " + binfo.Author + ", " +
		"Rating: " + strconv.Itoa(binfo.Rating) + "}"
}
