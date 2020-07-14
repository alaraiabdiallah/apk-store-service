package handlers

import (
	"github.com/alaraiabdiallah/apk-store-service/app"
	"github.com/alaraiabdiallah/apk-store-service/helpers"
	"github.com/alaraiabdiallah/apk-store-service/models"
	"github.com/labstack/echo"
	"net/http"
)

func SaveVersionHandler(c echo.Context) error {
	var v models.VersionDS
	if err := c.Bind(&v); err != nil { return err }
	if err := app.SaveVersion(v); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedJsonMessage(err))
	}
	return c.JSON(http.StatusOK, helpers.SuccessJsonMessage("Successfully to save version"))
}

func GetAllVersionHandler(c echo.Context) error {
	var results []models.VersionDS
	query := echo.Map{}

	for k, v := range c.QueryParams() {
		query[k] = v[0]
	}

	if err := app.GetAllVersion(query, &results); err != nil {return err}
	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"message": "",
		"data": results,
	})
}
