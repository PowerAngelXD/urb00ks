package service

import (
	"B00k/dao"
	"B00k/logger"
	"B00k/model"
	"errors"
)

type iUserInterface interface {
	IsUserExist(target int64) bool
	IsUserExistByName(target string) bool
	GetUser(target string) (model.UserInfo, error)
	GetUserById(target int64) (model.UserInfo, error)
	AddUser(name string, birth string, password string) error
	DeleteUser(target int64) error
	AddFav(target int64, title string) error
	DeleteFav(target_user string, target_book string) error
	UpdatePassword(target int64, pswd string) error
	UpdateName(target int64, name string) error
	UpdateAll(target int64, name string, pswd string, favs []string) error
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
		logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		logger.ServiceLog("Successfully create a user!")
		return nil
	}
}

// 删除用户
func (us *userService) DeleteUser(target int64) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to delete a unknown user in the library")
	}

	result := us.DB.Delete(target)

	if result != nil {
		logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		logger.ServiceLog("Successfully delete a user!")
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
		logger.ServiceLog("library error: occurred some problem, details: " + err.Error())
		return err
	} else {
		user.Favs = append(user.Favs, book)
		result := us.DB.UpdateAll(target, user.Name, user.Password, user.Favs)

		if result != nil {
			logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
			return err
		}

		logger.ServiceLog("Successfully create a user!")
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
		logger.ServiceLog("library error: occurred some problem, details: " + err.Error())
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
			logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
			return err
		}

		logger.ServiceLog("Successfully create a user!")
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
		logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		logger.ServiceLog("Successfully update a user!")
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
		logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		logger.ServiceLog("Successfully update a user!")
		return nil
	}
}

// 更新用户几乎所有字段
func (us *userService) UpdateAll(target int64, name string, pswd string, favs []string) error {
	if !us.IsUserExist(target) {
		return errors.New("library error: try to manage an unknown user")
	}

	result := us.DB.UpdateAll(target, name, pswd, favs)

	if result != nil {
		logger.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		logger.ServiceLog("Successfully update a user!")
		return nil
	}
}
