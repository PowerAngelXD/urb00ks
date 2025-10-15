package service

import (
	"B00k/middleware"

	"gorm.io/gorm"
)

type B00kService struct {
	LibSv  libService
	UserSv userService

	db *gorm.DB
}

var Service B00kService

func (service *B00kService) Init() {
	middleware.ServiceLog("Initializing the service...")
	Service.LibSv.DB.Link(service.db)
	middleware.ServiceLog("LibService: Connect to the database... Done!")
	Service.UserSv.DB.Link(service.db)
	middleware.ServiceLog("UserService: Connect to the database... Done!")
	middleware.ServiceLog("Service Initialization Done!")
}
