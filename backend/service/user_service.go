package service

import (
	"B00k/dao"
	"B00k/middleware"
	"B00k/model"
	"errors"
)

type iUserInterface interface {
	IsUserExist(target int64) bool
	IsUserExistByName(target string) bool
	GetUser(target string) (model.UserInfo, error)
	GetUserById(target int64) (model.UserInfo, error)
	AddUser(name string, birth string, password string) error
	AddFav(target int64, title string) error
	DeleteFav(target_user string, target_book string) error
	UpdatePassword(target int64, pswd string) error
	UpdateName(target int64, name string) error
}

type userService struct {
	DB dao.UserDao
}

// 查找某个用户是否存在
func (us *userService) IsUserExist(target int64) bool {
	return us.DB.IsExist(target)
}

func (us *userService) IsUserExistByName(target string) bool {
	return us.DB.IsExistByName(target)
}

// 获得一个UserInfo
func (us *userService) GetUser(target string) (model.UserInfo, error) {
	return us.DB.GetByName(target)
}

func (us *userService) GetUserById(target int64) (model.UserInfo, error) {
	return us.DB.Get(target)
}

// 添加一位用户
func (us *userService) AddUser(name string, birth string, password string) error {
	if us.IsUserExistByName(name) {
		return errors.New("library error: an existed user is already in the library")
	}

	result := us.DB.Create(name, birth, password)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully create a user!")
		return nil
	}
}

// 对指定用户增加喜好书籍
func (us *userService) AddFav(target int64, book string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: an existed user is already in the library")
	}

	user, err := us.DB.Get(target)

	if err != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + err.Error())
		return err
	} else {
		user.Favs = append(user.Favs, book)
		result := us.DB.UpdateAll(target, user.Name, user.Password, user.Favs)

		if result != nil {
			middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
			return err
		}

		middleware.ServiceLog("Successfully create a user!")
		return nil
	}
}

// 删除指定用户的喜好书籍
func (us *userService) DeleteFav(target int64, book string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	user, err := us.DB.Get(target)

	if err != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + err.Error())
		return err
	} else {
		new_favs := make([]string, len(user.Favs))
		for _, fav := range user.Favs {
			if fav != book {
				new_favs = append(new_favs, fav)
			}
		}

		result := us.DB.UpdateAll(target, user.Name, user.Password, new_favs)

		if result != nil {
			middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
			return err
		}

		middleware.ServiceLog("Successfully create a user!")
		return nil
	}
}

// 更新用户密码
func (us *userService) UpdatePassword(target int64, pswd string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	result := us.DB.UpdatePassword(target, pswd)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a user!")
		return nil
	}
}

// 更新用户名称
func (us *userService) UpdateName(target int64, name string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	result := us.DB.UpdateName(target, name)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a user!")
		return nil
	}
}
