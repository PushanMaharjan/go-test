package services

import (
	"go-fx-test/lib"
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
	return users, s.repository.Preload("Role").Find(&users).Error
}

func (s UserService) Create(user models.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) GetOneUser(userID lib.BinaryUUID) (user models.User, err error) {
	return user, s.repository.Preload("Role").First(&user, "id = ?", userID).Error
}

func (s UserService) UpdateUser(user models.User) error {
	if err := s.repository.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"fname":  user.Fname,
		"lname":  user.Lname,
		"roleID": user.RoleID,
	}).Error; err != nil {
		return err
	}
	return nil
}
