package controllers

import (
	"go-fx-test/lib"
	"go-fx-test/models"
	"go-fx-test/services"
	"go-fx-test/utils"
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

func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	user, err := u.service.GetOneUser(lib.ParseUUID(paramID))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

func (u UserController) UpdateUser(c *gin.Context) {

	userID := c.Param("id")

	var input utils.UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := u.service.GetOneUser(lib.ParseUUID(userID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"user not found": err.Error(),
		})
		return
	}

	if input.Fname != "" {
		user.Fname = input.Fname
	}

	if input.Lname != "" {
		user.Lname = input.Lname
	}

	if input.Admin != user.Admin {
		user.Admin = input.Admin
	}

	err = u.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user updated",
	})

}
