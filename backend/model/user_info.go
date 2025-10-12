package model

type UserInfo struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Birth    string   `json:"birth"`
	Password string   `json:"password"`
	Favs     []string `json:"favs"`
}
