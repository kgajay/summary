package controllers

import (
	"github.com/labstack/echo"
	"logger"
	"es"
	"net/http"
)

func EsIndices(c echo.Context) (err error) {

	esClient := es.GetESClient()
	names, err := esClient.IndexNames()
	if err != nil {
		logger.Log.Errorf("GET EsIndices:: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logger.Log.Infof("GET EsIndices indices:: %s", names)
	return c.JSON(http.StatusOK, names)
}
