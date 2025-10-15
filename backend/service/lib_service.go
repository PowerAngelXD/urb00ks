package service

import (
	"B00k/dao"
	"B00k/middleware"
	"B00k/model"
	"errors"
	"strconv"
)

type ilibService interface {
	IsBookExist(target int64) bool
	IsBookExistByTitle(target string) bool
	GetBook(target string) (model.BookInfo, error)
	GetBookById(target int64) (model.BookInfo, error)
	AddBook(name string, author string) error
	DeleteBook(target int64) error
	AddRecommend(title string, author string) error
	DeleteRecommend(title string) error
	UpdateBook(target int64, title string, author string, cvrurl string, rt int, views int) error
	UpdateCoverage(cvrurl string) error
	UpdateRating(rt int) error
	UpdateTitle(title string) error
	UpdateAuthor(author string) error
}

type libService struct {
	DB         *dao.LibDao
	Recommends []model.BookInfo
}

// 查找是否存在某本书
func (ls *libService) IsBookExist(target int64) bool {
	return ls.DB.IsExist(target)
}

func (ls *libService) IsBookExistByTitle(target string) bool {
	return ls.DB.IsExistByTitle(target)
}

// 获得一个BookInfo
func (ls *libService) GetBook(target string) (model.BookInfo, error) {
	return ls.DB.GetByTitle(target)
}

func (ls *libService) GetBookById(target int64) (model.BookInfo, error) {
	return ls.DB.Get(target)
}

// 增加库中书本
func (ls *libService) AddBook(title string, author string) error {
	return ls.DB.Create(title, author, "/img/Book.png")
}

// 删除库中书本
func (ls *libService) DeleteBook(target int64) error {
	return ls.DB.Delete(target)
}

// 增加推荐书本
func (ls *libService) AddRecommend(target int64, author string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	book, err := ls.DB.Get(target)

	if err != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + err.Error())
		return err
	} else {
		ls.Recommends = append(ls.Recommends, book)
		middleware.ServiceLog("Successfully add a recommend!")
		return nil
	}
}

// 删除推荐书本
func (ls *libService) DeleteRecommend(target int64) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to delete an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	_, err := ls.DB.Get(target)

	if err != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + err.Error())
		return err
	} else {
		recommends := make([]model.BookInfo, len(ls.Recommends))
		for _, book := range ls.Recommends {
			if book.Id != target {
				recommends = append(recommends, book)
			}
		}
		ls.Recommends = recommends
		middleware.ServiceLog("Successfully delete a recommend!")
		return nil
	}
}

// 更新书籍状态
func (ls *libService) UpdateBook(target int64, title string, author string, cvrurl string, rating int, views int) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	result := ls.DB.UpdateAll(target, title, author, rating, views)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a book!")
		return nil
	}
}

func (ls *libService) UpdateCoverage(target int64, cvrurl string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	result := ls.DB.UpdateUrl(target, cvrurl)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a book!")
		return nil
	}
}

func (ls *libService) UpdateRating(target int64, rt int) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	result := ls.DB.UpdateRating(target, rt)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a book!")
		return nil
	}
}

func (ls *libService) UpdateTitle(target int64, title string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	result := ls.DB.UpdateTitle(target, title)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a book!")
		return nil
	}
}

func (ls *libService) UpdateAuthor(target int64, author string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	result := ls.DB.UpdateAuthor(target, author)

	if result != nil {
		middleware.ServiceLog("library error: occurred some problem, details: " + result.Error())
		return result
	} else {
		middleware.ServiceLog("Successfully update a book!")
		return nil
	}
}
