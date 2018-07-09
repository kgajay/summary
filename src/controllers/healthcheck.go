package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"time"
	"dao"
)

var startTime time.Time

func GetStatus(c echo.Context) (err error) {
	resp := map[string]interface{}{
		"status": "up",
		"uptime": uptime(),
	}
	return c.JSON(http.StatusOK, resp)
}

// GetDeepStatus provides rest endpoint to check the status of the external db factory dependency
func GetDeepStatus(c echo.Context) (err error) {
	testdb := dao.GetDb()
	err = testdb.DB().Ping()
	resp := make(map[string]interface{})
	if err != nil {
		resp["db"] = map[string]interface{}{
			"status": "down",
			"error":  err.Error(),
		}
	} else {
		resp["db"] = map[string]interface{}{
			"status": "up",
		}
	}
	return c.JSON(http.StatusOK, resp)
}

func init() {
	startTime = time.Now()
}

func uptime() string {
	return time.Since(startTime).String()
}
