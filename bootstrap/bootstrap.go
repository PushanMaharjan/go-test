package bootstrap

import (
	"context"
	"go-fx-test/api/controllers"
	"go-fx-test/api/routes"
	"go-fx-test/infrastructure"
	"go-fx-test/lib"
	"go-fx-test/repository"
	"go-fx-test/services"
	"log"

	"go.uber.org/fx"
)

var Module = fx.Options(
	routes.Module,
	infrastructure.Module,
	lib.Module,
	controllers.Module,
	services.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle,
	router infrastructure.Router,
	routes routes.Routes,
	database infrastructure.Database) {
	conn, _ := database.DB.DB()
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Println("application started")
				conn.SetMaxOpenConns(10)
				routes.Setup()
				router.Run(":5000")
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("application stopped")
			_ = conn.Close()
			return nil
		},
	})
}
