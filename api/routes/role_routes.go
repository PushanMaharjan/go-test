package routes

import (
	"go-fx-test/api/controllers"
	"go-fx-test/infrastructure"
	"log"
)

// UserRoutes struct
type RoleRoutes struct {
	handler        infrastructure.Router
	roleController controllers.RoleController
}

func NewRoleRoutes(
	handler infrastructure.Router,
	roleController controllers.RoleController,
) RoleRoutes {
	return RoleRoutes{
		handler:        handler,
		roleController: roleController,
	}
}

// Setup user routes
func (s RoleRoutes) Setup() {
	log.Println("Setting up routes")
	api := s.handler.Group("/api")
	{
		api.GET("/roles", s.roleController.GetRole)
		api.POST("/roles", s.roleController.SaveRole)
		api.GET("/roles/:id", s.roleController.GetOneRole)
		api.PUT("/roles/:id", s.roleController.UpdateRole)
	}
}
