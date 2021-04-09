package mongods

import (
	"time"

	"github.com/alaraiabdiallah/apk-store-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var versionCollName string = "version"

func upsertVersionParams(filter interface{}, data models.VersionDS, update_data *bson.D) {
	var version models.VersionDS
	_ = FindVersion(filter, &version)

	created_at := time.Now()
	if version.CreatedAt.IsZero() {
		created_at = version.CreatedAt
	}

	set_data := bson.M{
		"flag":       data.Flag,
		"version":    data.Version,
		"build_code": data.BuildCode,
		"created_at": created_at,
		"updated_at": time.Now(),
	}
	*update_data = bson.D{{"$set", set_data}}
}

func UpsertVersion(data models.VersionDS) error {
	opts := options.Update().SetUpsert(true)
	coll, dbCtx := collection(versionCollName)
	defer dbCtx.Disconnect()
	filter := bson.D{{"flag", data.Flag}}
	var update_data bson.D
	upsertVersionParams(filter, data, &update_data)
	if _, err := coll.UpdateOne(dbCtx.ctx, filter, update_data, opts); err != nil {
		return err
	}
	return nil
}

func FindAllVersion(filter interface{}, results *[]models.VersionDS) error {
	coll, dbCtx := collection(versionCollName)
	defer dbCtx.Disconnect()
	cur, err := coll.Find(dbCtx.ctx, filter)
	if err != nil {
		return err
	}
	for cur.Next(dbCtx.ctx) {
		var res models.VersionDS
		err := cur.Decode(&res)
		if err != nil {
			return err
		}
		*results = append(*results, res)
	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}

func FindVersion(filter interface{}, result *models.VersionDS) error {
	coll, dbCtx := collection(versionCollName)
	defer dbCtx.Disconnect()
	if err := coll.FindOne(dbCtx.ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}
