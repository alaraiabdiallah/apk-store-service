package app

import (
	"bytes"
	"errors"
	"fmt"
	ds "github.com/alaraiabdiallah/apk-store-service/data_sources/mongods"
	"github.com/alaraiabdiallah/apk-store-service/models"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

func getFileExt(filename string) string{
	s := strings.Split(filename, ".")
	return s[len(s)-1]
}

func generatedFilename(filename string) string{
	ext := getFileExt(filename)
	return fmt.Sprintf("%v.%v", primitive.NewObjectID().Hex(),ext)
}

func getFileMime(file *multipart.FileHeader) (types.Type, error){
	buf := bytes.NewBuffer(nil)
	src, err := file.Open()
	if err != nil { return types.Type{}, err }
	defer src.Close()
	if _, err := io.Copy(buf, src); err != nil { return types.Type{}, err }
	kind, err := filetype.Match(buf.Bytes())
	if err != nil { return types.Type{}, err }
	return kind, nil;
}

func saveToStorage(file *multipart.FileHeader, file_data *models.MediaDS) error{
	file_name := generatedFilename(file.Filename)
	file_ext := getFileExt(file.Filename)
	file_path := "storage/"+ file_name
	file_mime := "application/vnd.android.package-archive"
	dst, err := os.Create(file_path)
	if err != nil { return err }
	defer dst.Close()
	src, err := file.Open()
	if err != nil { return err }
	defer src.Close()
	if file_ext != "apk" { return errors.New("Extension file must be APK") }
	if _, err = io.Copy(dst, src); err != nil { return err }
	*file_data = models.MediaDS{
		Id:       primitive.NewObjectID(),
		Filename: file_name,
		Filepath: file_path,
		Mime:     file_mime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return nil
}

func SaveMedia(params models.MediaUploadParams, result *models.MediaDS) error {
	var file_data models.MediaDS
	file := params.File
	if err := saveToStorage(file, &file_data); err != nil{
		log.Fatal(err)
		return err
	}

	file_data.Flag = params.Flag
	file_data.Version = params.Version

	if err := ds.CreateMedia(file_data); err != nil { return err }

	*result = file_data
	return nil
}

func GetAllMedia(filter models.MediaFilter, results *interface{}) error {
	if(filter.OnlyLink){
		var medias []models.MediaLink
		if err := ds.FindAllMediaLink(filter.Query, &medias); err != nil { return err }
		*results = medias
	} else {
		var medias []models.MediaDS
		if err := ds.FindAllMedia(filter.Query, &medias); err != nil { return err }
		*results = medias
	}

	return nil
}

func GetOneMedia(filter map[string]interface{}, result *models.MediaDS) error {
	if err := ds.FindMedia(filter, result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("Not Found")
		}
		return err
	}

	return nil
}