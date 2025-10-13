package service

import (
	"B00k/model"
	"errors"
	"slices"
)

type iUserInterface interface {
	IsUserExist(target string)
	GetUser(target string)
	GetUserById(target int)
	AddUser(name string, birth string, password string)
	AddFav(target string, title string)
	DeleteFav(target_user string, target_book string)
	UpdatePassword(target string, pswd string)
	UpdateName(target string, name string)
}

type userService struct {
	LibObject *model.B0Lib
}

// 查找某个用户是否存在
func (us *userService) IsUserExist(target string) bool {
	_, ok := us.LibObject.Users[target]
	return ok
}

// 获得一个UserInfo
func (us *userService) GetUser(target string) (model.UserInfo, error) {
	if us.IsUserExist(target) {
		return us.LibObject.Users[target], nil
	} else {
		return model.UserInfo{Name: "[NIL]"}, errors.New("library error: try to get a unknown book")
	}
}

func (us *userService) GetUserFromId(target int) (model.UserInfo, error) {
	for _, user := range us.LibObject.Users {
		if user.Id == target {
			return user, nil
		}
	}
	return model.UserInfo{Id: -1, Name: "Unkn0wn User"}, errors.New("library error: try to get a unknown book")
}

// 添加一位用户
func (us *userService) AddUser(name string, birth string, password string) error {
	if us.IsUserExist(name) {
		return errors.New("library error: an existed user is already in the library")
	}

	return nil
}

// 对指定用户增加喜好书籍
func (us *userService) AddFav(target string, book string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: an existed user is already in the library")
	}

	copy := us.LibObject.Users[target]
	copy.Favs = append(copy.Favs, book)
	us.LibObject.Users[target] = copy

	return nil
}

// 删除指定用户的喜好书籍
func (us *userService) DeleteFav(target string, book string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	tmp_user := us.LibObject.Users[target]
	if slices.Contains(tmp_user.Favs, book) {
		favs := make([]string, len(tmp_user.Favs)-1)
		for _, name := range tmp_user.Favs {
			if name != book {
				favs = append(favs, name)
			}
		}
		tmp_user.Favs = favs
		us.LibObject.Users[target] = tmp_user

		return nil
	}

	return errors.New("library error: try to delete a unknown favorite book")
}

// 更新用户密码
func (us *userService) UpdatePassword(target string, pswd string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	tmp_user := us.LibObject.Users[target]

	if tmp_user.Password == pswd {
		return errors.New("library error: try to update a same password")
	}

	tmp_user.Password = pswd
	us.LibObject.Users[target] = tmp_user
	return nil
}

// 更新用户名称
func (us *userService) UpdateName(target string, name string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	tmp_user := us.LibObject.Users[target]

	if tmp_user.Name == name {
		return errors.New("library error: try to update a same password")
	}

	tmp_user.Name = name
	us.LibObject.Users[target] = tmp_user
	return nil
}
