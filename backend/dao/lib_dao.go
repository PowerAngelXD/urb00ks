package dao

import (
	"B00k/logger"
	"B00k/model"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type iLibDao interface {
	Count() int64
	Link(target_db *gorm.DB)
	IsExistByTitle(target string) bool
	IsExist(target int64) bool
	Create(title string, author string, url string) error
	UpdateAll(target int64, title string, author string, url string, rating int, views int) error
	UpdateTitle(target int64, title string) error
	UpdateAuthor(target int64, author string) error
	UpdateRating(target int64, rt int) error
	UpdateViews(target int64, views int) error
	UpdateUrl(target int64, url string) error
	Get(target int64) (model.BookInfo, error)
	GetByTitle(target string) (model.BookInfo, error)
	GetInNum(num int) ([]model.BookInfo, error)
	GetAll() ([]model.BookInfo, error)
	Delete(target int64) error
}

type LibDao struct {
	DB *gorm.DB
}

func (ld *LibDao) Count() (int64, error) {
	var max int64
	result := ld.DB.Model(&model.BookInfo{}).Count(&max)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot get the count of 'book_list', details: " + result.Error.Error())
		return -1, result.Error
	} else {
		logger.DBLog("Operation: successfully get the count in the table: 'book_list'")
		return max, nil
	}
}

func (ld *LibDao) Link(target_db *gorm.DB) {
	ld.DB = target_db
	logger.DBLog("Database linked done!")
}

func (ld *LibDao) IsExistByTitle(target string) bool {
	var book model.BookInfo
	result := ld.DB.Where("title = ?", target).First(&book)

	if result.Error != nil {
		return false
	} else {
		return true
	} // TODO: error可能不是未找到key，之后做适配处理
}

func (ld *LibDao) IsExist(target int64) bool {
	var book model.BookInfo
	result := ld.DB.First(&book, target)
	if result.Error != nil {
		return false
	} else {
		return true
	} // TODO: error可能不是未找到key，之后做适配处理
}

func (ld *LibDao) Create(title string, author string, url string) error {
	result := ld.DB.Create(&model.BookInfo{Title: title, Author: author, ImageUrl: url})

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot create the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: create an book... Done")
		return nil
	}
}

func (ld *LibDao) UpdateAll(target int64, title string, url string, author string, rating int, views int) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	result := ld.DB.Model(&book).Updates(map[string]any{
		"title":     title,
		"author":    author,
		"rating":    rating,
		"views":     views,
		"image_url": url,
	})

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot update the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully update the book: [" + strconv.FormatInt(target, 10) + "] all the fields")
		return nil
	}
}

func (ld *LibDao) UpdateTitle(target int64, title string) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	result := ld.DB.Model(&book).Update("title", title)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot update the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully update the book: [" + strconv.FormatInt(target, 10) + "] target field: Title")
		return nil
	}
}

func (ld *LibDao) UpdateAuthor(target int64, author string) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	result := ld.DB.Model(&book).Update("author", author)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot update the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully update the book: [" + strconv.FormatInt(target, 10) + "] target field: Author")
		return nil
	}
}

func (ld *LibDao) UpdateRating(target int64, rt int) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	result := ld.DB.Model(&book).Update("rating", rt)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot update the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully update the book: [" + strconv.FormatInt(target, 10) + "] target field: Rating")
		return nil
	}
}

func (ld *LibDao) UpdateViews(target int64, views int) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	result := ld.DB.Model(&book).Update("views", views)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot update the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully update the book: [" + strconv.FormatInt(target, 10) + "] target field: Title")
		return nil
	}
}

func (ld *LibDao) UpdateUrl(target int64, url string) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot update a unknown book")
	}

	var book model.BookInfo
	book.Id = target
	result := ld.DB.Model(&book).Update("image_url", url)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot update the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully update the book: [" + strconv.FormatInt(target, 10) + "] target field: ImageUrl")
		return nil
	}
}

func (ld *LibDao) Get(target int64) (model.BookInfo, error) {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return model.BookInfo{}, errors.New("cannot get a unknown book")
	}

	var book model.BookInfo
	result := ld.DB.First(&book, target)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot get the book instance, details: " + result.Error.Error())
		return model.BookInfo{}, result.Error
	} else {
		logger.DBLog("Operation: successfully get the book: [" + strconv.FormatInt(target, 10) + "]")
		return book, nil
	}
}

func (ld *LibDao) GetByTitle(target string) (model.BookInfo, error) {
	if !ld.IsExistByTitle(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return model.BookInfo{}, errors.New("cannot get a unknown book")
	}

	var book model.BookInfo
	result := ld.DB.Where("title = ?", target).First(&book)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot get the book instance, details: " + result.Error.Error())
		return model.BookInfo{}, result.Error
	} else {
		logger.DBLog("Operation: successfully get the book: [" + target + "]")
		return book, nil
	}
}

func (ld *LibDao) GetInNum(num int, offset int) ([]model.BookInfo, error) {
	var books []model.BookInfo
	logger.DBLog("offset: " + strconv.Itoa(offset))
	result := ld.DB.Limit(int(num)).Offset(offset).Find(&books)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot get the book instance, details: " + result.Error.Error())
		return nil, result.Error
	} else {
		return books, nil
	}
}

func (ld *LibDao) GetAll() ([]model.BookInfo, error) {
	var books []model.BookInfo
	result := ld.DB.Find(&books)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot get the book instance, details: " + result.Error.Error())
		return nil, result.Error
	} else {
		logger.DBLog("Successfully get all the book!")
		return books, nil
	}
}

func (ld *LibDao) Delete(target int64) error {
	if !ld.IsExist(target) {
		logger.DBLog("Occurred error: Cannot update a unknown book")
		return errors.New("cannot delete a unknown book")
	}

	result := ld.DB.Delete(&model.BookInfo{}, target)

	if result.Error != nil {
		logger.DBLog("Occurred error: Cannot delete the book instance, details: " + result.Error.Error())
		return result.Error
	} else {
		logger.DBLog("Operation: successfully delete the book: [" + strconv.FormatInt(target, 10) + "]")
		return nil
	}
}
