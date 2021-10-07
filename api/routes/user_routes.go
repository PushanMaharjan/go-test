package routes

import (
	"go-fx-test/api/controllers"
	"go-fx-test/infrastructure"
	"log"
)

// UserRoutes struct
type UserRoutes struct {
	handler        infrastructure.Router
	userController controllers.UserController
}

func NewUserRoutes(
	handler infrastructure.Router,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	log.Println("Setting up routes")
	api := s.handler.Group("/api")
	{
		api.GET("/user", s.userController.GetUser)
		api.POST("/user", s.userController.SaveUser)
		api.GET("/user/:id", s.userController.GetOneUser)
		api.PUT("/user/:id", s.userController.UpdateUser)
		api.POST("/user/send-email", s.userController.TriggerEmail)
	}
}
