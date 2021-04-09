package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alaraiabdiallah/apk-store-service/app"
	"github.com/alaraiabdiallah/apk-store-service/helpers"
	"github.com/alaraiabdiallah/apk-store-service/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMedia(c echo.Context) error {
	var results interface{}
	only_link := false
	query := echo.Map{}

	//if par := c.QueryParam("flag"); par != "" { query["flag"] = par}
	for k, v := range c.QueryParams() {
		query_par_val := v[0]
		if k == "only-link" && query_par_val == "true" {
			only_link = true
		} else {
			query[k] = query_par_val
		}
	}
	filter := models.MediaFilter{OnlyLink: only_link, Query: query}
	if err := app.GetAllMedia(filter, &results); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": "",
		"data":    results,
	})
}

func ShowMediaFile(c echo.Context) error {
	var result models.MediaDS
	media_id := c.Param("media_id")
	id, _ := primitive.ObjectIDFromHex(media_id)
	filter := echo.Map{"_id": id}
	if err := app.GetOneMedia(filter, &result); err != nil {
		if err.Error() == "Not Found" {
			c.JSON(http.StatusNotFound, helpers.FailedJsonMessage(err.Error()))
		}
		return err
	}
	f, err := os.Open(result.Filepath)
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "application/vnd.android.package-archive", f)
}

func UploadMedia(c echo.Context) error {
	file, err := c.FormFile("file")
	var file_data models.MediaDS
	file_params := models.MediaUploadParams{
		Flag:      c.FormValue("flag"),
		Version:   c.FormValue("version"),
		File:      file,
		BuildCode: c.FormValue("build_code"),
	}
	if file == nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedJsonMessage("File body not defined"))
	}
	if file_params.Flag == "" {
		return c.JSON(http.StatusBadRequest, helpers.FailedJsonMessage("Flag body empty or not defined"))
	}
	if file_params.Version == "" {
		return c.JSON(http.StatusBadRequest, helpers.FailedJsonMessage("Version body empty or not defined"))
	}
	if file_params.BuildCode == "" {
		return c.JSON(http.StatusBadRequest, helpers.FailedJsonMessage("Build code body empty or not defined"))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedJsonMessage(err))
	}
	if err := app.SaveMedia(file_params, &file_data); err != nil {
		errCode := http.StatusInternalServerError
		if err.Error() == "build_code invalid." {
			errCode = http.StatusBadRequest
		}
		return c.JSON(errCode, helpers.FailedJsonMessage(err.Error()))

	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": "Successfully to upload media",
		"data":    file_data.Id,
	})
}

func GetLatestVersionAPKByFlag(c echo.Context) error {
	var result models.MediaDS
	flag := c.QueryParam("flag")
	if flag == "" {
		c.JSON(http.StatusBadRequest, helpers.FailedJsonMessage("Flag parameter not defined"))
	}
	filter := echo.Map{"flag": flag}
	if err := app.GetOneMedia(filter, &result); err != nil {
		if err.Error() == "Not Found" {
			c.JSON(http.StatusNotFound, helpers.FailedJsonMessage(err.Error()))
		}
		return err
	}
	f, err := os.Open(result.Filepath)
	if err != nil {
		return err
	}
	fmt.Printf("Version: %v, Flag: %v", result.Version, result.Flag)
	return c.Stream(http.StatusOK, "application/vnd.android.package-archive", f)
}
