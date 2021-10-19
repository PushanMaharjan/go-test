package services

import (
	"fmt"
	"go-fx-test/lib"
	"go-fx-test/models"
	"go-fx-test/repository"
	"go-fx-test/utils/email_templates"
	"log"
)

type UserService struct {
	repository          repository.UserRepository
	notificationService NotificationService
}

func NewUserService(
	repository repository.UserRepository,
	notificationService NotificationService,
) UserService {
	return UserService{
		repository:          repository,
		notificationService: notificationService,
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

func (s UserService) TriggerTestEmailToUser(
	username string, email string,
) error {
	type EmailData struct {
		Username string
	}

	emailData := EmailData{
		Username: username,
	}

	fmt.Println("from user input", "email:", email, "username:", username)

	err := s.notificationService.WithTemplate(email, email_templates.TestEmail+"/", email_templates.TestEmail, emailData)
	if err != nil {
		log.Println("Error sending email: ", err)
	}
	return err
}
