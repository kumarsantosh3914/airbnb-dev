package app

import (
	dbConfig "ReviewService/config/db"
	config "ReviewService/config/env"
	"ReviewService/controllers"
	repo "ReviewService/db/repositories"
	"ReviewService/router"
	"ReviewService/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // PORT
}

type Application struct {
	Config Config
}

// Construction for config
func NewConfig() Config {
	port := config.GetString("PORT", ":3004")

	return Config{
		Addr: port,
	}
}

// Construction for Application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error setting up database: ", err)
		return err
	}

	rr := repo.NewReviewRepository(db)
	rs := services.NewReviewService(rr)
	rc := controllers.NewReviewController(rs)
	rRouter := router.NewReviewRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(rRouter),
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
		IdleTimeout:  10 * time.Second, // Set idle timeout to 10 seconds
	}

	fmt.Println("Review Service is running on port", app.Config.Addr)

	return server.ListenAndServe()
}
