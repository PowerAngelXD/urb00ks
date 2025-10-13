package service

import (
	"B00k/middleware"
	"B00k/model"
)

type b00kService struct {
	LibObject *model.B0Lib
	LibSv     libService
	UserSv    userService
}

var Service b00kService

func init() {
	model.Lib.Init()
	Service.LibObject = &model.Lib
	middleware.ServiceLog("Initialize the service's library object... Done")
	Service.LibSv.LibObject = &model.Lib
	middleware.ServiceLog("Linking the library object to the libService... Done")
	Service.UserSv.LibObject = &model.Lib
	middleware.ServiceLog("Linking the library object to the UserService... Done")
	middleware.ServiceLog("Service Initialization Done!")
}
