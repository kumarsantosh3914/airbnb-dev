package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // PORT
}

type Application struct {
	Config Config
	// Store  db.Storage
}

// Construction for Config
func NewConfig() Config {
	port := config.GetString("PORT", ":8080")

	return Config{
		Addr: port,
	}
}

// Constuction for Application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		// Store:  *db.NewStorage(),
	}
}

func (app *Application) Run() error {

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error setting up database: ", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	rpr := repo.NewRolePermissionRepository(db)
	urr := repo.NewUserRoleRepository(db)
	us := services.NewUserService(ur)
	rs := services.NewRoleService(rr, rpr, urr)
	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	uRouter := router.NewUserRouter(uc)
	rRouter := router.NewRoleRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter, rRouter),
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
		IdleTimeout:  10 * time.Second, // Set idle timeout to 10 seconds
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}
