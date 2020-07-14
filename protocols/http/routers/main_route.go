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
	v1 := e.Group("/v1")
	apk := v1.Group("/apk")
	apk.GET("/", handlers.GetAllMedia, middlewares.APIAuth)
	apk.GET("/get-latest-version", handlers.GetLatestVersionAPKByFlag, middlewares.APIAuthWithQuery)
	apk.GET("/:media_id", handlers.ShowMediaFile)
	apk.POST("/", handlers.UploadMedia, middlewares.APIAuth)

	ver := v1.Group("/versions")
	ver.GET("/", handlers.GetAllVersionHandler, middlewares.APIAuth)
	ver.PATCH("/", handlers.SaveVersionHandler, middlewares.APIAuth)

}
