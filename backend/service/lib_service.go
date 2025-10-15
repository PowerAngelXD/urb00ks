package service

import (
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
	AddBookByInstance(book model.BookInfo) error
	DeleteBook(target int64) error
	AddRecommend(title string, author string) error
	DeleteRecommend(title string) error
	UpdateBook(title string, author string, cvrurl string, rt int) error
	UpdateCoverage(cvrurl string) error
	UpdateRating(rt int) error
	UpdateTitle(title string) error
	UpdateAuthor(author string) error
}

type libService struct {
	LibObject *model.B0Lib
}

// 查找是否存在某本书
func (ls *libService) IsBookExist(target int64) bool {
	_, ok := ls.LibObject.Books[target]
	return ok
}

func (ls *libService) IsBookExistByTitle(target string) bool {
	for _, book := range ls.LibObject.Books {
		if book.Title == target {
			return true
		}
	}
	return false
}

// 获得一个BookInfo
func (ls *libService) GetBook(target string) (model.BookInfo, error) {
	if ls.IsBookExistByTitle(target) {
		for _, user := range ls.LibObject.Books {
			if user.Title == target {
				return user, nil
			}
		}
		return model.BookInfo{}, errors.New("library error: try to get a unknown book")
	} else {
		return model.BookInfo{}, errors.New("library error: try to get a unknown book")
	}
}

func (ls *libService) GetBookById(target int64) (model.BookInfo, error) {
	for _, book := range ls.LibObject.Books {
		if book.Id == target {
			return book, nil
		}
	}
	return model.BookInfo{Id: -1, Title: "Unkn0wn B0ok"}, errors.New("library error: try to get a unknown book")
}

// 增加库中书本
func (ls *libService) AddBook(title string, author string) error {
	if ls.IsBookExistByTitle(title) {
		return errors.New("library error: an existed book is already in the library")
	}

	return nil
}

// 增加库中书本，直接传入一个BookInfo进行增加操作
func (ls *libService) AddBookByInstance(book model.BookInfo) error {
	if ls.IsBookExistByTitle(book.Title) {
		return errors.New("library error: an existed book is already in the library")
	}

	ls.LibObject.Books[book.Id] = book

	return nil
}

// 删除库中书本
func (ls *libService) DeleteBook(target int64) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to delete an unknown book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	delete(ls.LibObject.Books, target)

	return nil
}

// 增加推荐书本
func (ls *libService) AddRecommend(title string, author string) error {
	if ls.IsBookExistByTitle(title) {
		return errors.New("library error: an existed recommended book is already in the library")
	}

	// TODO: Dao层要实现对Recommend的GetId方法

	return nil
}

// 删除推荐书本
func (ls *libService) DeleteRecommend(target int64) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to delete an unknown recommended book: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	// TODO: Dao层要实现对Recommend的GetId方法

	return nil
}

// 更新书籍状态
func (ls *libService) UpdateBook(title string, author string, cvrurl string, rating int) error {
	for idx, book := range ls.LibObject.Books {
		if book.Title == title {
			updatedBook := model.BookInfo{Id: book.Id, Title: title, Author: author, Rating: rating, ImageUrl: cvrurl}
			ls.LibObject.Books[idx] = updatedBook

			return nil
		}
	}

	return errors.New("library error: want to update an unknown book: \"" + title + "\"")
}

func (ls *libService) UpdateCoverage(target int64, cvrurl string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.ImageUrl = cvrurl
	ls.LibObject.Books[target] = copy

	return nil
}

func (ls *libService) UpdateRating(target int64, rt int) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.Rating = rt
	ls.LibObject.Books[target] = copy

	return nil
}

func (ls *libService) UpdateTitle(target int64, title string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.Title = title
	ls.LibObject.Books[target] = copy

	return nil
}

func (ls *libService) UpdateAuthor(target int64, author string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + strconv.FormatInt(target, 10) + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.Author = author
	ls.LibObject.Books[target] = copy

	return nil
}
