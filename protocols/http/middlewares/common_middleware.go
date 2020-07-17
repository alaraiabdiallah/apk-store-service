package middlewares

import (
	"errors"
	"github.com/alaraiabdiallah/apk-store-service/helpers"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func APIAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		api_key := c.Request().Header.Get("apikey")
		if api_key != os.Getenv("API_KEY") {
			err := errors.New("x-api-key invalid")
			return c.JSON(http.StatusForbidden, helpers.FailedJsonMessage(err.Error()))
		}
		return next(c)
	}
}

func APIAuthWithQuery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		api_key := c.QueryParam("apikey")
		if api_key != os.Getenv("API_KEY") {
			err := errors.New("apikey invalid")
			return c.JSON(http.StatusForbidden, helpers.FailedJsonMessage(err.Error()))
		}
		return next(c)
	}
}
