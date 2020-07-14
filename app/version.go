package app

import (
	"errors"
	ds "github.com/alaraiabdiallah/apk-store-service/data_sources/mongods"
	"github.com/alaraiabdiallah/apk-store-service/models"
)

func SaveVersion(params models.VersionDS) error {
	if err := ds.UpsertVersion(params); err != nil { return err }
	return nil
}

func GetAllVersion(filter map[string]interface{}, results *[]models.VersionDS) error {
	if err := ds.FindAllVersion(filter, results); err != nil { return err }
	return nil
}

func GetOneVersion(filter map[string]interface{}, result *models.VersionDS) error {
	if err := ds.FindVersion(filter, result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("Not Found")
		}
		return err
	}

	return nil
}