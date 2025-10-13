package dao

import (
	"B00k/middleware"
	"B00k/model"
	"errors"

	"gorm.io/gorm"
)

type iLibDao interface {
	IsExistByTitle(target string) bool
	IsExist(target int) bool
	Create(title string, author string) error
	UpdateAll(target int, title string, author string, rating int, views int) error
	UpdateTitle(target int, title string) error
	UpdateAuthor(target int, author string) error
	UpdateRating(target int, rt int) error
	UpdateViews(target int, views int) error
	Get(target int) (model.BookInfo, error)
	Delete(target int) error
}

type LibDao struct {
	db *gorm.DB
}

func (ld *LibDao) IsExistByTitle(target string) bool {
	var book model.BookInfo
	result := ld.db.Where("title = ?", target).First(&book)

	if result.Error != nil {
		return false
	} else {
		return true
	} // TODO: error可能不是未找到key，之后做适配处理
}

func (ld *LibDao) IsExist(target int) bool {
	var book model.BookInfo
	result := ld.db.First(&book, target)
	if result.Error != nil {
		return false
	} else {
		return true
	} // TODO: error可能不是未找到key，之后做适配处理
}

func (ld *LibDao) Create(title string, author string) error {
	result := ld.db.Create(model.BookInfo{Title: title, Author: author})

	if result.Error != nil {
		middleware.DBLog("Occurred error: Cannot create the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		return nil
	}
}

func (ld *LibDao) UpdateAll(target int, title string, author string, rating int, views int) error {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	ld.db.Model(&book).Updates(map[string]any{
		"title":  title,
		"author": author,
		"rating": rating,
		"views":  views,
	})

	return nil
}

func (ld *LibDao) UpdateTitle(target int, title string) error {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	ld.db.Model(&book).Update("title", title)

	return nil
}

func (ld *LibDao) UpdateAuthor(target int, author string) error {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	ld.db.Model(&book).Update("author", author)

	return nil
}

func (ld *LibDao) UpdateRating(target int, rt int) error {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	ld.db.Model(&book).Update("rating", rt)

	return nil
}

func (ld *LibDao) UpdateViews(target int, views int) error {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	ld.db.Model(&book).Update("views", views)

	return nil
}

func (ld *LibDao) Get(target int) (model.BookInfo, error) {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return model.BookInfo{Id: -1, Title: "Unkn0wn B0ok"}, errors.New("cannot get a unknown book")
	}

	var book model.BookInfo
	ld.db.First(&book, target)

	return book, nil
}

func (ld *LibDao) Delete(target int) error {
	if !ld.IsExist(target) {
		middleware.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot delete a unknown book")
	}

	ld.db.Delete(&model.BookInfo{}, target)

	return nil
}
