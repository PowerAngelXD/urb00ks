package model

import "strconv"

type UserInfo struct {
	Id       int64    `json:"id" gorm:"primaryKey;type:int;column:id;autoIncrement"`
	Name     string   `json:"name" gorm:"not null;type:varchar(100);column:name"`
	Birth    string   `json:"birth" gorm:"not null;type:varchar(50);column:birth"`
	Password string   `json:"password" gorm:"not null;type:varchar(20);column:password"`
	Favs     []string `json:"favs" gorm:"type:text;serializer:json;column:favs"`
}

func (UserInfo) TableName() string {
	return "user_list"
}

func (uinfo *UserInfo) ToString() string {
	result := "{Book: \"" + strconv.FormatInt(uinfo.Id, 10) + ": " + uinfo.Name + "\", " +
		"Birth: " + uinfo.Birth +
		"Favs: {"
	for _, fav := range uinfo.Favs {
		result += fav + ", "
	}
	result += "}"
	return result
}
