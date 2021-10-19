package controllers

import (
	"go-fx-test/lib"
	"go-fx-test/models"
	"go-fx-test/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service     services.UserService
	roleService services.RoleService
}

func NewUserController(userService services.UserService, roleService services.RoleService) UserController {
	return UserController{
		service:     userService,
		roleService: roleService,
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

type userResponse struct {
	user models.User
	role models.Role
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

	var userRes userResponse
	userRes.user = user

	c.JSON(200, gin.H{
		"data": user,
	})

}

func (u UserController) UpdateUser(c *gin.Context) {

	userID := c.Param("id")

	var input models.UpdateUserInput

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

	if input.RoleID != "" {
		user.RoleID = lib.ParseUUID(input.RoleID)
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

func (u UserController) TriggerEmail(c *gin.Context) {
	var input models.EmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(c, http.StatusBadRequest, err)
		return
	}
	err := u.service.TriggerTestEmailToUser(input.Username, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "message sent",
	})

}
