package mongods

import (
	"context"
	"errors"
	"fmt"
	"github.com/alaraiabdiallah/apk-store-service/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type MongoDSContext struct {
	ctx context.Context
	db *mongo.Database
	client *mongo.Client
}

type MongoDSOptions struct {
	timeout time.Duration
}

func (ctx *MongoDSContext) Ping() error{
	return ctx.client.Ping(ctx.ctx, readpref.Primary())
}

func (ctx *MongoDSContext) Disconnect(){
	ctx.client.Disconnect(ctx.ctx)
}

func CheckConnection() {
	mctx := Ctx()
	defer mctx.Disconnect()
	if err := mctx.Ping(); err != nil {
		if err.Error() == "context deadline exceeded" {
			err = errors.New("Failed to connect mongo")
		}
		log.Fatal(err)
	}else{
		log.Println("Success to connect mongo")
	}
}

func collection(collName string) (*mongo.Collection,*MongoDSContext) {
	ctx := Ctx()
	return ctx.db.Collection(collName), ctx
}

func uriConnBuilder() string{
	cred := ""
	host := helpers.EnvGetByDefault("MONGO_HOST","localhost")
	port := helpers.EnvGetByDefault("MONGO_PORT","27017")
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	if user != "" && pass != "" {
		cred = user + ":" + pass
	}
	return fmt.Sprintf("mongodb://%v@%v:%v",cred,host,port)
}

func defineOptions(opts ...MongoDSOptions) MongoDSOptions{
	mopts := MongoDSOptions{timeout:2*time.Second}
	if len(opts) > 0 {
		mopts = opts[0]
	}
	return mopts
}

func Ctx(opts ...MongoDSOptions) *MongoDSContext{
	mopts := defineOptions(opts...)
	ctx, _ := context.WithTimeout(context.Background(), mopts.timeout)
	uri := uriConnBuilder()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	dbname := helpers.EnvGetByDefault("MONGO_DBNAME","db-media")
	db := client.Database(dbname)
	return &MongoDSContext{ctx: ctx, db: db, client: client}
}