package mongods

import (
	"fmt"
	"github.com/alaraiabdiallah/apk-store-service/helpers"
	"github.com/alaraiabdiallah/apk-store-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var mediaCollname string = "media"

func getMediaUrl(media_id interface{}) string {
	default_port := os.Getenv("HTTP_PROT_PORT")
	base_url := helpers.EnvGetByDefault("APP_BASE_URL","http://localhost:"+default_port)
	id := media_id.(primitive.ObjectID).Hex()
	return fmt.Sprintf("%v/v1/apk/%v",base_url, id)
}

func CreateMedia(media interface{}) error{
	coll, dbCtx := collection(mediaCollname)
	defer dbCtx.Disconnect()
	if _, err := coll.InsertOne(dbCtx.ctx, media); err != nil {
		return err
	}
	return nil
}

func FindAllMedia(filter interface{}, results *[]models.MediaDS) error {
	coll, dbCtx := collection(mediaCollname)
	defer dbCtx.Disconnect()
	sort_opts := options.Find()
	sort_opts.SetSort(bson.D{{"version", -1}})
	cur, err := coll.Find(dbCtx.ctx, filter, sort_opts)
	if err != nil { return err }
	for cur.Next(dbCtx.ctx) {
		var res models.MediaDS
		err := cur.Decode(&res)
		if err != nil { return err }
		*results = append(*results, res)
	}
	if err := cur.Err(); err != nil { return err }
	return nil
}

func FindAllMediaLink(filter interface{}, results *[]models.MediaLink) error {
	var media []models.MediaDS
	if err := FindAllMedia(filter, &media); err != nil{return err}
	for _, res := range media {
		*results = append(*results, models.MediaLink{Url: getMediaUrl(res.Id), Version: res.Version, Flag: res.Flag})
	}
	return nil
}

func FindMedia(filter interface{}, result *models.MediaDS) error {
	coll, dbCtx := collection(mediaCollname)
	defer dbCtx.Disconnect()
	sort_opts := options.FindOne()
	sort_opts.SetSort(bson.D{{"version", -1}})
	if err := coll.FindOne(dbCtx.ctx,filter, sort_opts).Decode(result); err != nil { return err }
	return nil
}