package services

import (
	"go-fx-test/models"
	"go-fx-test/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(
	repository repository.UserRepository,
) UserService {
	return UserService{
		repository: repository,
	}
}

func (s UserService) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

func (s UserService) Create(user models.User) error {
	return s.repository.Create(&user).Error
}
