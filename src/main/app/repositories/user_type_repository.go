package repositories

import (
	"service-api/src/main/app/entities"

	"gorm.io/gorm"
)

type UserTypeRepo interface {
	FindAll() ([]entities.UserType, error)
	FindById(id uint) (entities.UserType, error)
	FindByType(userType string) (entities.UserType, error)
	Save(userType entities.UserType) (entities.UserType, error)
}

type userTypeRepo struct {
	db *gorm.DB
}
