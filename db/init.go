package db

import (
	"context"
	"fmt"
	"time"

	"github.com/letanthang/my_framework/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName = "my_user"
)

var Client *mongo.Client

func init() {
	// c, _ := mongo.NewClient(options.Client().SetAuth(options.Credential{}))

	var uri string
	uri = "mongodb://" + config.Config.MongoDB.Username + ":" + config.Config.MongoDB.Password + "@" + config.Config.MongoDB.Host + ":27017"
	fmt.Println(uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
}
