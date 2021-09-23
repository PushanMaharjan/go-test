package controllers

import (
	"go-fx-test/models"
	"go-fx-test/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		service: userService,
	}
}

func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{"data": users})
}

func (u UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.Create(user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}
