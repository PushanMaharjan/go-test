package bootstrap

import (
	"context"
	"log"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("start")
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("stop")
			return nil
		},
	})
}
