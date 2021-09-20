package main

import (
	"go-fx-test/bundlefx"
	"go-fx-test/httphandler"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.Module,
		fx.Invoke(httphandler.HandlerFunc),
	).Run()
}
