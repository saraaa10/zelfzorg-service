package repositories

import (
	"service-api/src/main/app/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindById(id uint) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByUsername(username string) (entities.User, error)
	Save(user entities.User) (entities.User, error)
}

type userRepo struct {
	db *gorm.DB
}
