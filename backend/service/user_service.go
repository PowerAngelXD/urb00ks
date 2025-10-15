package service

import (
	"B00k/model"
	"errors"
	"slices"
)

type iUserInterface interface {
	IsUserExist(target int64) bool
	IsUserExistByName(target string) bool
	GetUser(target string) (model.UserInfo, error)
	GetUserById(target int64) (model.UserInfo, error)
	AddUser(name string, birth string, password string) error
	AddUserByInstance(user model.UserInfo) error
	AddFav(target int64, title string) error
	DeleteFav(target_user string, target_book string) error
	UpdatePassword(target int64, pswd string) error
	UpdateName(target int64, name string) error
}

type userService struct {
	LibObject *model.B0Lib
}

// 查找某个用户是否存在
func (us *userService) IsUserExist(target int64) bool {
	_, ok := us.LibObject.Users[target]
	return ok
}

func (ls *userService) IsUserExistByName(target string) bool {
	for _, book := range ls.LibObject.Users {
		if book.Name == target {
			return true
		}
	}
	return false
}

// 获得一个UserInfo
func (us *userService) GetUser(target string) (model.UserInfo, error) {
	if us.IsUserExistByName(target) {
		for _, user := range us.LibObject.Users {
			if user.Name == target {
				return user, nil
			}
		}
		return model.UserInfo{}, errors.New("library error: try to get a unknown user")
	} else {
		return model.UserInfo{}, errors.New("library error: try to get a unknown user")
	}
}

func (us *userService) GetUserById(target int64) (model.UserInfo, error) {
	for _, user := range us.LibObject.Users {
		if user.Id == target {
			return user, nil
		}
	}
	return model.UserInfo{Id: -1, Name: "Unkn0wn User"}, errors.New("library error: try to get a unknown user")
}

// 添加一位用户
func (us *userService) AddUser(name string, birth string, password string) error {
	if us.IsUserExistByName(name) {
		return errors.New("library error: an existed user is already in the library")
	}

	return nil
}

func (us *userService) AddUserByInstance(user model.UserInfo) error {
	if us.IsUserExistByName(user.Name) {
		return errors.New("library error: an existed user is already in the library")
	}

	us.LibObject.Users[user.Id] = user

	return nil
}

// 对指定用户增加喜好书籍
func (us *userService) AddFav(target int64, book string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: an existed user is already in the library")
	}

	copy := us.LibObject.Users[target]
	copy.Favs = append(copy.Favs, book)
	us.LibObject.Users[target] = copy

	return nil
}

// 删除指定用户的喜好书籍
func (us *userService) DeleteFav(target int64, book string) error {
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
func (us *userService) UpdatePassword(target int64, pswd string) error {
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
func (us *userService) UpdateName(target int64, name string) error {
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
