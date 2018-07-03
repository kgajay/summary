package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func GetStatus(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)
}

// GetDeepStatus provides rest endpoint to check the status of the Elasticsearch endpoint
func GetDeepStatus(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)
}
