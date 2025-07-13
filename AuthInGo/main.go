package main

import (
	"AuthInGo/app"
)

func main() {
	cfg := app.NewConfig()
	app := app.Application{
		Config: cfg,
	}

	app.Run()
}
