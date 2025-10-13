package model

type UserInfo struct {
	Id       int      `json:"id" gorm:"primaryKey;type:int;column:id"`
	Name     string   `json:"name" gorm:"not null;type:varchar(100);column:name"`
	Birth    string   `json:"birth" gorm:"not null;type:varchar(50);column:birth"`
	Password string   `json:"password" gorm:"not null;type:varchar(20);column:password"`
	Favs     []string `json:"favs" gorm:"type:text;serializer:json;column:favs"`
}
