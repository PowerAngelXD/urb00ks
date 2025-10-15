package service

import (
	"B00k/dao"
	"B00k/middleware"
	"B00k/model"
)

type b00kService struct {
	LibObject *model.B0Lib
	LibSv     libService
	UserSv    userService

	libDB  dao.LibDao
	userDB dao.UserDao
}

var Service b00kService

func init() {
	model.Lib.Init()
	Service.LibObject = &model.Lib
	middleware.ServiceLog("Initialize the service's library object... Done")
	Service.LibSv.LibObject = &model.Lib
	Service.libDB.Link(dao.OfficialDB)
	count, err := Service.libDB.Count()
	if err != nil {
		middleware.ServiceLog("Failed to get the count of the table: 'book_list', details: " + err.Error())
		panic("Failed to get the count of the table: 'book_list', details: " + err.Error())
	} else {
		var i int64 = 0
		for ; i < count; i++ {
			book, err := Service.libDB.Get(i)
			if err != nil {
				middleware.ServiceLog("Failed to add the book to the table: 'book_list', so you might miss a book, details: " + err.Error())
			} else {
				Service.LibSv.AddBookByInstance(book)
			}
		}
	}
	middleware.ServiceLog("Linking the library object to the libService... Done")
	Service.UserSv.LibObject = &model.Lib
	Service.userDB.Link(dao.OfficialDB)
	if err != nil {
		middleware.ServiceLog("Failed to get the count of the table: 'user_list', details: " + err.Error())
		panic("Failed to get the count of the table: 'user_list', details: " + err.Error())
	} else {
		var i int64 = 0
		for ; i < count; i++ {
			user, err := Service.userDB.Get(i)
			if err != nil {
				middleware.ServiceLog("Failed to add the book to the table: 'user_list', so you might miss a user, details: " + err.Error())
			} else {
				Service.UserSv.AddUserByInstance(user)
			}
		}
	}
	middleware.ServiceLog("Linking the library object to the UserService... Done")
	middleware.ServiceLog("Service Initialization Done!")
}
