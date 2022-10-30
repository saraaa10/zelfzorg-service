package usecases

import "service-api/src/main/entities"

type UserTypeRepo interface {
	FindAll() ([]entities.UserType, error)
	FindById(id uint) (entities.UserType, error)
	FindByType(userType string) (entities.UserType, error)
	Save(userType entities.UserType) (entities.UserType, error)
}
