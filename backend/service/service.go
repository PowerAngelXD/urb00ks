package service

import (
	"B00k/logger"

	"gorm.io/gorm"
)

type B00kService struct {
	LibSv  libService
	UserSv userService

	db *gorm.DB
}

var Service B00kService

func (service *B00kService) Init() {
	logger.ServiceLog("Initializing the service...")
	Service.LibSv.DB.Link(service.db)
	logger.ServiceLog("LibService: Connect to the database... Done!")
	Service.UserSv.DB.Link(service.db)
	logger.ServiceLog("UserService: Connect to the database... Done!")
	logger.ServiceLog("Service Initialization Done!")
}
