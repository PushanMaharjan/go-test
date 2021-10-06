package controllers

import (
	"go-fx-test/lib"
	"go-fx-test/models"
	"go-fx-test/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return RoleController{
		service: roleService,
	}
}

func (r RoleController) GetRole(c *gin.Context) {
	roles, err := r.service.GetAllRoles()
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{"data": roles})
}

func (r RoleController) SaveRole(c *gin.Context) {
	role := models.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := r.service.Create(role); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "role created"})
}

func (r RoleController) GetOneRole(c *gin.Context) {
	paramID := c.Param("id")

	role, err := r.service.GetOneRole(lib.ParseUUID(paramID))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"role not found": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": role,
	})

}

func (r RoleController) UpdateRole(c *gin.Context) {

	roleID := c.Param("id")

	var input models.UpdateRoleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	role, err := r.service.GetOneRole(lib.ParseUUID(roleID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"role not found": err.Error(),
		})
		return
	}

	if input.Role != "" {
		role.Role = input.Role
	}

	err = r.service.UpdateRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "role updated",
	})

}
