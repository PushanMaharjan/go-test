package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"go-fx-test/bootstrap"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
