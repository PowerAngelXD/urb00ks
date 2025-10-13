package service

import (
	"B00k/model"
	"errors"
)

type ilibService interface {
	IsBookExist(title string)
	GetBook(target string)
	GetBookById(target int)
	AddBook(name string, author string)
	AddBookByInstance(book model.BookInfo)
	DeleteBook(target string)
	AddRecommend(title string, author string)
	DeleteRecommend(title string)
	UpdateBook(title string, author string, cvrurl string, rt int)
	UpdateCoverage(cvrurl string)
	UpdateRating(rt int)
	UpdateTitle(title string)
	UpdateAuthor(author string)
}

type libService struct {
	LibObject *model.B0Lib
}

// 查找是否存在某本书
func (ls *libService) IsBookExist(title string) bool {
	_, ok := ls.LibObject.Books[title]
	return ok
}

// 获得一个BookInfo
func (ls *libService) GetBook(target string) (model.BookInfo, error) {
	if ls.IsBookExist(target) {
		return ls.LibObject.Books[target], nil
	} else {
		return model.BookInfo{Id: -1, Title: "Unkn0wn B0ok"}, errors.New("library error: try to get a unknown book")
	}
}

func (ls *libService) GetBookById(target int) (model.BookInfo, error) {
	for _, book := range ls.LibObject.Books {
		if book.Id == target {
			return book, nil
		}
	}
	return model.BookInfo{Id: -1, Title: "Unkn0wn B0ok"}, errors.New("library error: try to get a unknown book")
}

// 增加库中书本
func (ls *libService) AddBook(title string, author string) error {
	if ls.IsBookExist(title) {
		return errors.New("library error: an existed book is already in the library")
	}

	return nil
}

// 增加库中书本，直接传入一个BookInfo进行增加操作
func (ls *libService) AddBookByInstance(book model.BookInfo) error {
	if ls.IsBookExist(book.Title) {
		return errors.New("library error: an existed book is already in the library")
	}

	ls.LibObject.Books[book.Title] = book

	return nil
}

// 删除库中书本
func (ls *libService) DeleteBook(target string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to delete an unknown book: \"" + target + "\"")
	}

	delete(ls.LibObject.Books, target)

	return nil
}

// 增加推荐书本
func (ls *libService) AddRecommend(title string, author string) error {
	if ls.IsBookExist(title) {
		return errors.New("library error: an existed recommended book is already in the library")
	}

	// TODO: Dao层要实现GetId方法

	return nil
}

// 删除推荐书本
func (ls *libService) DeleteRecommend(target string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to delete an unknown recommended book: \"" + target + "\"")
	}

	// TODO: Dao层要实现GetId方法

	return nil
}

// 更新书籍状态
func (ls *libService) UpdateBook(title string, author string, cvrurl string, rating int) error {
	for idx, book := range ls.LibObject.Books {
		if book.Title == title {
			updatedBook := model.BookInfo{Id: book.Id, Title: title, Author: author, Rating: rating, CoverUrl: cvrurl}
			ls.LibObject.Books[idx] = updatedBook

			return nil
		}
	}

	return errors.New("library error: want to update an unknown book: \"" + title + "\"")
}

func (ls *libService) UpdateCoverage(target string, cvrurl string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + target + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.CoverUrl = cvrurl
	ls.LibObject.Books[target] = copy

	return nil
}

func (ls *libService) UpdateRating(target string, rt int) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + target + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.Rating = rt
	ls.LibObject.Books[target] = copy

	return nil
}

func (ls *libService) UpdateTitle(target string, title string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + target + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.Title = title
	ls.LibObject.Books[target] = copy

	return nil
}

func (ls *libService) UpdateAuthor(target string, author string) error {
	if !ls.IsBookExist(target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + target + "\"")
	}

	copy := ls.LibObject.Books[target]
	copy.Author = author
	ls.LibObject.Books[target] = copy

	return nil
}
