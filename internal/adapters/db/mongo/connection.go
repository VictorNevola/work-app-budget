package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type (
	Connection struct {
		*mongo.Database
	}
)

func NewMongoDB() (*Connection, func()) {
	uri := "mongodb://root:example@localhost:27017/"
	mainDatabase := "work-app-budget"

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic("Failed to connect to Mongo \n Error: " + err.Error())
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		panic("Failed to connect to Mongo \n Error: " + err.Error())
	}

	disconnectFunction := func() {
		client.Disconnect(context.TODO())
	}

	return &Connection{
		Database: client.Database(mainDatabase),
	}, disconnectFunction
}
