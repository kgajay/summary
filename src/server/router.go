package server

import (
	"controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter for routing requests
func NewRouter() *echo.Echo {
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.RequestID())

	// Endpoints for healthcheck
	router.GET("/status", controllers.GetStatus)
	router.GET("/deepstatus", controllers.GetDeepStatus)

	// RestAPI for User
	router.GET("/user/:id", controllers.GetUser)
	router.GET("/user/show", controllers.ShowUser)
	router.POST("/user", controllers.SaveUser)

	return router
}
