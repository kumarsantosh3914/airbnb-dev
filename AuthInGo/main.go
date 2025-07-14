package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
)

func main() {
	config.Load()

	cfg := app.NewConfig()
	app := app.Application{
		Config: cfg,
	}

	dbConfig.SetupDB()

	app.Run()
}
