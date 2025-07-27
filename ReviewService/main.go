package main

import (
	"ReviewService/app"
	dbConfig "ReviewService/config/db"
	config "ReviewService/config/env"
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
