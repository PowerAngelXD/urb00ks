package api

import "B00k/model"

type ReturnStruct[T model.BookInfo | model.UserInfo | int | []model.BookInfo | []model.UserInfo] struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
}
