package router

import (
	env "AuthInGo/config/env"
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	// chiRouter.Use(middlewares.RequestLogger) // Use the RequestLogger middleware

	chiRouter.Use(middleware.Logger) // Built-in Chi middleware for logging requests

	chiRouter.Use(middlewares.RateLimitMiddleware) // Apply rate limiting middleware

	chiRouter.Get("/ping", controllers.PingHandler)

	chiRouter.HandleFunc("/hotelservice/*", utils.ProxyToService(env.GetString("HOTEL_SERVICE_URL", ""), "/hotelservice"))
	chiRouter.HandleFunc("/bookingservice/*", utils.ProxyToService(env.GetString("BOOKING_SERVICE_URL", ""), "/bookingservice"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
}
