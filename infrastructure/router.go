package infrastructure

import (
	"go-fx-test/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	env lib.Env,
) Router {
	httpRouter := gin.Default()

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	return Router{
		httpRouter,
	}
}
