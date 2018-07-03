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

	// Endpoints for healthcheck
	router.GET("/status", controllers.GetStatus)
	router.GET("/deepstatus", controllers.GetDeepStatus)

	return router
}
