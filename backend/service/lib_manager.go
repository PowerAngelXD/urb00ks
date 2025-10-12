package service

import (
	"B00k/model"
	"errors"
)

var bookCode int = 0 // 用于标识书本BookCode的变量

// 书籍管理 ======================
// 用于生成书籍id
func ConsumeBookId() int {
	bookCode++
	return bookCode
}

// 查找是否存在某本书
func IsBookExist(lib *model.B0Lib, name string) bool {
	_, ok := lib.Books[name]
	return ok
}

// 获得一个BookInfo
func GetBook(lib *model.B0Lib, target string) (model.BookInfo, error) {
	if IsBookExist(lib, target) {
		return lib.Books[target], nil
	} else {
		return model.BookInfo{Id: -1, Title: "Unkn0wn B0ok"}, errors.New("library error: try to get a unknown book")
	}
}

// 增加库中书本
func AddBook(lib *model.B0Lib, name string, author string) error {
	if IsBookExist(lib, name) {
		return errors.New("library error: an existed book is already in the library")
	}

	lib.Books[name] = model.BookInfo{Id: ConsumeBookId(), Title: name, Author: author, Rating: 0, CoverUrl: "null"}

	return nil
}

// 增加库中书本，直接传入一个BookInfo进行增加操作
func AddBookInstance(lib *model.B0Lib, book model.BookInfo) error {
	if IsBookExist(lib, book.Title) {
		return errors.New("library error: an existed book is already in the library")
	}

	lib.Books[book.Title] = book

	return nil
}

// 删除库中书本
func DeleteBook(lib *model.B0Lib, name string) error {
	if !IsBookExist(lib, name) {
		return errors.New("library error: want to delete an unknown book: \"" + name + "\"")
	}

	delete(lib.Books, name)

	return nil
}

// 增加推荐书本
func AddRecommend(lib *model.B0Lib, name string, author string) error {
	if IsBookExist(lib, name) {
		return errors.New("library error: an existed recommended book is already in the library")
	}

	lib.Recommends[name] = model.BookInfo{Id: ConsumeBookId(), Title: name, Author: author, Rating: 0, CoverUrl: "null"}

	return nil
}

// 删除推荐书本
func DeleteRecommend(lib *model.B0Lib, name string) error {
	if !IsBookExist(lib, name) {
		return errors.New("library error: want to delete an unknown recommended book: \"" + name + "\"")
	}

	delete(lib.Recommends, name)

	return nil
}

// 更新书籍状态
func UpdateBook(lib *model.B0Lib, name string, author string, cvrurl string, rating int) error {
	for idx, book := range lib.Books {
		if book.Title == name {
			updatedBook := model.BookInfo{Id: book.Id, Title: name, Author: author, Rating: rating, CoverUrl: cvrurl}
			lib.Books[idx] = updatedBook

			return nil
		}
	}

	return errors.New("library error: want to update an unknown book: \"" + name + "\"")
}

func UpdateCoverage(lib *model.B0Lib, target string, cvrurl string) error {
	if !IsBookExist(lib, target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + target + "\"")
	}

	copy := lib.Books[target]
	copy.CoverUrl = cvrurl
	lib.Books[target] = copy

	return nil
}

func UpdateRating(lib *model.B0Lib, target string, rt int) error {
	if !IsBookExist(lib, target) {
		return errors.New("library error: want to update an unknown book for coverage: \"" + target + "\"")
	}

	copy := lib.Books[target]
	copy.Rating = rt
	lib.Books[target] = copy

	return nil
}
