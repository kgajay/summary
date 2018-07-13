package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"model"
	"services"
	"logger"
	"strconv"
)

// e.GET("/user/:id", getUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	iid, errs := strconv.Atoi(id)
	if errs != nil {
		return c.NoContent(http.StatusNotFound)
	}

	status, resp := services.GetUser(iid)
	logger.Log.Infof("GetUser status: %s, resp: %s", status, resp)
	return c.JSON(status, resp)
}

// e.GET("/user/show", show)
func ShowUser(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/user", save)
func SaveUser(c echo.Context) (err error) {

	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	status, resp := services.CreateUser(u)
	logger.Log.Infof("SaveUser status: %s, resp: %s", status, resp)
	return c.JSON(status, resp)
}
