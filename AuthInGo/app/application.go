package app

import (
	"AuthInGo/clients"
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

	// Initialize notification client
	notificationServiceURL := config.GetString("NOTIFICATION_SERVICE_URL", "http://localhost:3002")
	notificationClient := clients.NewNotificationClient(notificationServiceURL)

	ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	rpr := repo.NewRolePermissionRepository(db)
	urr := repo.NewUserRoleRepository(db)
	or := repo.NewOTPRepository(db)
	rs := services.NewRoleService(rr, rpr, urr)
	us := services.NewUserService(ur, rs)
	os := services.NewOTPService(or, notificationClient)
	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	oc := controllers.NewOTPController(os)
	uRouter := router.NewUserRouter(uc)
	rRouter := router.NewRoleRouter(rc)
	oRouter := router.NewOTPRouter(oc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter, rRouter, oRouter),
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
		IdleTimeout:  10 * time.Second, // Set idle timeout to 10 seconds
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}
