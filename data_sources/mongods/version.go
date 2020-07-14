package mongods

import (
	"github.com/alaraiabdiallah/apk-store-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var versionCollName string = "version"

func UpsertVersion(data models.VersionDS) error{
	opts := options.Update().SetUpsert(true)
	coll, dbCtx := collection(versionCollName)
	defer dbCtx.Disconnect()
	bdata := bson.M{
		"flag": data.Flag,
		"version": data.Version,
	}
	filter := bson.D{{"flag", data.Flag}}
	update := bson.D{{"$set", bdata}}
	if _, err := coll.UpdateOne(dbCtx.ctx, filter, update, opts); err != nil {
		return err
	}
	return nil
}

func FindAllVersion(filter interface{}, results *[]models.VersionDS) error {
	coll, dbCtx := collection(versionCollName)
	defer dbCtx.Disconnect()
	cur, err := coll.Find(dbCtx.ctx, filter)
	if err != nil { return err }
	for cur.Next(dbCtx.ctx) {
		var res models.VersionDS
		err := cur.Decode(&res)
		if err != nil { return err }
		*results = append(*results, res)
	}
	if err := cur.Err(); err != nil { return err }
	return nil
}

func FindVersion(filter interface{}, result *models.VersionDS) error {
	coll, dbCtx := collection(versionCollName)
	defer dbCtx.Disconnect()
	if err := coll.FindOne(dbCtx.ctx,filter).Decode(result); err != nil { return err }
	return nil
}