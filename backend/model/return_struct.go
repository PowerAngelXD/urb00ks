package model

type ReturnStruct[T string | BookInfo | UserInfo | int | []BookInfo | []UserInfo] struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
}
