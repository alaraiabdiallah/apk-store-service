package routers

import (
	"fmt"
	"github.com/alaraiabdiallah/apk-store-service/app"
	"github.com/alaraiabdiallah/apk-store-service/protocols/http/handlers"
	"github.com/alaraiabdiallah/apk-store-service/protocols/http/middlewares"
	"github.com/labstack/echo"
	"net/http"
)

func HttpRouters(e *echo.Echo){
	e.GET("/", func(c echo.Context) error {
		version := fmt.Sprintf("App version %v", app.Ctx().GetVersion())
		return c.String(http.StatusOK, version)
	})
	v1 := e.Group("/v1/apk")
	v1.GET("/", handlers.GetAllMedia, middlewares.APIAuth)
	v1.GET("/get-latest-version", handlers.GetLatestVersionByFlag, middlewares.APIAuthWithQuery)
	v1.GET("/:media_id", handlers.ShowMediaFile)
	v1.POST("/", handlers.UploadMedia, middlewares.APIAuth)
}
