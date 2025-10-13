package service

import (
	"B00k/model"
	"fmt"
)

type b00kService struct {
	LibObject model.B0Lib
	LibSv     libService
	UserSv    userService
}

var Service b00kService

func init() {
	Service.LibObject = model.Lib
	Service.LibObject.Init() // 初始化Library
	fmt.Println("[Service] Initialize the service's library object... Done")
	Service.LibSv.LibObject = &model.Lib
	fmt.Println("[Service] Linking the library object to the libService... Done")
	Service.UserSv.LibObject = &model.Lib
	fmt.Println("[Service] Linking the library object to the UserService... Done")
	fmt.Println("[Service] Service Initialization Done!")
}
