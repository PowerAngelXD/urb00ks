package service

import (
	"B00k/model"
	"errors"
	"slices"
)

var userCode = 0

// 生成用户ID并返回
func ConsumeUserId() int {
	userCode++
	return userCode
}

// 查找某个用户是否存在
func IsUserExist(lib *model.B0Lib, target string) bool {
	_, ok := lib.Users[target]
	return ok
}

// 获得一个UserInfo
func GetUser(lib *model.B0Lib, target string) (model.UserInfo, error) {
	if IsUserExist(lib, target) {
		return lib.Users[target], nil
	} else {
		return model.UserInfo{Name: "[NIL]"}, errors.New("library error: try to get a unknown book")
	}
}

// 添加一位用户
func AddUser(lib *model.B0Lib, name string, birth string, password string) error {
	if IsUserExist(lib, name) {
		return errors.New("library error: an existed user is already in the library")
	}

	lib.Users[name] = model.UserInfo{Id: ConsumeUserId(), Name: name, Birth: birth, Password: password}

	return nil
}

// 对指定用户增加喜好书籍
func AddFav(lib *model.B0Lib, user string, book string) error {
	if !IsUserExist(lib, user) {
		return errors.New("library error: an existed user is already in the library")
	}

	copy := lib.Users[user]
	copy.Favs = append(copy.Favs, book)
	lib.Users[user] = copy

	return nil
}

// 删除指定用户的喜好书籍
func DeleteFav(lib *model.B0Lib, user string, book string) error {
	if !IsUserExist(lib, user) {
		return errors.New("library error: try to manage an unknown user")
	}

	tmp_user := lib.Users[user]
	if slices.Contains(tmp_user.Favs, book) {
		favs := make([]string, len(tmp_user.Favs)-1)
		for _, name := range tmp_user.Favs {
			if name != book {
				favs = append(favs, name)
			}
		}
		tmp_user.Favs = favs
		lib.Users[user] = tmp_user

		return nil
	}

	return errors.New("library error: try to delete a unknown favorite book")
}

// 更新用户密码
func UpdatePassword(lib *model.B0Lib, user string, pswd string) error {
	if !IsUserExist(lib, user) {
		return errors.New("library error: try to manage an unknown user")
	}

	tmp_user := lib.Users[user]

	if tmp_user.Password == pswd {
		return errors.New("library error: try to update a same password")
	}

	tmp_user.Password = pswd
	lib.Users[user] = tmp_user
	return nil
}
