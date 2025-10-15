package dao

import (
	"B00k/middleware"
	"B00k/model"
	"errors"
	"slices"
	"strconv"

	"gorm.io/gorm"
)

type iUserDao interface {
	Count() (int64, error)
	Link(target_db *gorm.DB)
	IsExistByName(target string) bool
	IsExist(target int64) bool
	Create(name string, birth string, pswd string) error
	UpdateAll(target int64, name string, pswd string, favs []string) error
	UpdateName(target int64, name string) error
	UpdatePassword(target int64, pswd string) error
	AddFav(target int64, fav string) error
	DeleteFav(target int64, fav string) error
	Get(target int64) (model.UserInfo, error)
	GetByName(target string) (model.UserInfo, error)
	Delete(target int64) error
}

type UserDao struct {
	db *gorm.DB
}

func (ld *UserDao) Count() (int64, error) {
	var max int64
	result := ld.db.Model(&model.UserInfo{}).Count(&max)

	if result.Error != nil {
		middleware.DBLog("Occurred error: Cannot get the count of 'user_list', details: " + result.Error.Error())
		return -1, result.Error
	} else {
		middleware.DBLog("Operation: successfully get the count in the table: 'user_list'")
		return max, nil
	}
}

func (ld *UserDao) Link(target_db *gorm.DB) {
	ld.db = target_db
	middleware.DBLog("User: Database linked done!")
}

func (ud *UserDao) IsExistByName(target string) bool {
	var user model.UserInfo
	result := ud.db.Where("name = ?", target).First(&user)

	if result.Error != nil {
		return false
	} else {
		return true
	} // TODO: error可能不是未找到key，之后做适配处理
}

func (ud *UserDao) IsExist(target int64) bool {
	var user model.UserInfo
	result := ud.db.First(&user, target)
	if result.Error != nil {
		return false
	} else {
		return true
	} // TODO: error可能不是未找到key，之后做适配处理
}

func (ud *UserDao) Create(name string, birth string, pswd string) error {
	result := ud.db.Create(model.UserInfo{Name: name, Birth: birth, Password: pswd})

	if result.Error != nil {
		middleware.DBLog("Occurred error: Cannot create the user instance, details: " + result.Error.Error())
		return result.Error
	} else {
		middleware.DBLog("Operation: create an user... Done")
		return nil
	}
}

func (ud *UserDao) UpdateAll(target int64, name string, pswd string, favs []string) error {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update an unknown user")
		return errors.New("cannot update an unknown user")
	}

	var user model.UserInfo
	user.Id = target
	result := ud.db.Model(&user).Updates(map[string]any{
		"name":     name,
		"password": pswd,
		"favs":     favs,
	})

	if result.Error != nil {
		middleware.DBLog("Occurred error: In operation UpdateAll: " + result.Error.Error())
		return result.Error
	} else {
		middleware.DBLog("Operation: successfully update the user: [" + strconv.FormatInt(target, 10) + "] all the fields")
		return nil
	}
}

func (ud *UserDao) UpdateName(target int64, name string) error {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update an unknown user")
		return errors.New("cannot update an unknown user")
	}

	var user model.UserInfo
	user.Id = target
	result := ud.db.Model(&user).Update("name", name)

	if result.Error != nil {
		middleware.DBLog("Occurred error: In operation UpdateName " + result.Error.Error())
		return result.Error
	} else {
		middleware.DBLog("Operation: successfully update the user: [" + strconv.FormatInt(target, 10) + "] target field: Name")
		return nil
	}
}

func (ud *UserDao) UpdatePassword(target int64, pswd string) error {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update an unknown user")
		return errors.New("cannot update an unknown user")
	}

	var user model.UserInfo
	user.Id = target
	result := ud.db.Model(&user).Update("password", pswd)

	if result.Error != nil {
		middleware.DBLog("Occurred error: In operation UpdatePassword " + result.Error.Error())
		return result.Error
	} else {
		middleware.DBLog("Operation: successfully update the user: [" + strconv.FormatInt(target, 10) + "] target field: Password")
		return nil
	}
}

func (ud *UserDao) AddFav(target int64, fav string) error {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update an unknown user")
		return errors.New("cannot update an unknown user")
	}

	var user model.UserInfo
	user.Id = target
	ud.db.First(&user, target)

	if slices.Contains(user.Favs, fav) {
		middleware.DBLog("Occurred error: Cannot add a new favorite for the user")
		return errors.New("cannot add a new favorite for the user")
	}

	user.Favs = append(user.Favs, fav)
	result := ud.db.Model(&user).Updates(user)

	if result.Error != nil {
		middleware.DBLog("Occurred error: In operation AddFavs " + result.Error.Error())
		return result.Error
	} else {
		middleware.DBLog("Operation: successfully update the user: [" + strconv.FormatInt(target, 10) + "] target field: Favs")
		return nil
	}
}

func (ud *UserDao) DeleteFav(target int64, fav string) error {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update an unknown user")
		return errors.New("cannot update an unknown user")
	}

	var user model.UserInfo
	user.Id = target
	ud.db.First(&user, target)

	if !slices.Contains(user.Favs, fav) {
		middleware.DBLog("Occurred error: Cannot delete a unknown favorite for the user")
		return errors.New("cannot delete a unknown favorite for the user")
	}

	user.Favs = append(user.Favs, fav)
	result := ud.db.Model(&user).Updates(user)

	if result.Error != nil {
		middleware.DBLog("Occurred error: In operation DeleteFavs " + result.Error.Error())
		return result.Error
	} else {
		middleware.DBLog("Operation: successfully update the user: [" + strconv.FormatInt(target, 10) + "] target field: Favs")
		return nil
	}
}

func (ud *UserDao) Get(target int64) (model.UserInfo, error) {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot get an unknown user")
		return model.UserInfo{}, errors.New("cannot get an unknown user")
	}

	var user model.UserInfo
	result := ud.db.First(&user, target)

	if result.Error != nil {
		middleware.DBLog("Occurred error: Cannot get the user, detail: " + result.Error.Error())
		return model.UserInfo{}, errors.New("cannot get the user, detail: " + result.Error.Error())
	} else {
		middleware.DBLog("Operation: successfully get the user: [" + strconv.FormatInt(target, 10) + "]")
		return user, nil
	}
}

func (ud *UserDao) GetByName(target string) (model.UserInfo, error) {
	if !ud.IsExistByName(target) {
		middleware.DBLog("Occurred error: Cannot get an unknown user")
		return model.UserInfo{}, errors.New("cannot get an unknown user")
	}

	var user model.UserInfo
	result := ud.db.Where("name = ?", target).First(&user)

	if result.Error != nil {
		middleware.DBLog("Occurred error: Cannot get the user, detail: " + result.Error.Error())
		return model.UserInfo{}, errors.New("cannot get the user, detail: " + result.Error.Error())
	} else {
		middleware.DBLog("Operation: successfully get the user: [" + target + "]")
		return user, nil
	}
}

func (ud *UserDao) Delete(target int64) error {
	if !ud.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot delete an unknown user")
		return errors.New("cannot delete an unknown user")
	}

	result := ud.db.Delete(model.UserInfo{}, target)

	if result.Error != nil {
		middleware.DBLog("Occurred error: Cannot delete the user, detail: " + result.Error.Error())
		return errors.New("cannot delete the user, detail: " + result.Error.Error())
	} else {
		middleware.DBLog("Operation: successfully delete the user: [" + strconv.FormatInt(target, 10) + "]")
		return nil
	}
}
