package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoStorage struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoStorage() *MongoStorage {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin_pg:*****@pgcluster.g4nlk.mongodb.net/parking_lot?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	db := client.Database("parking_lot")
	mongoStorage := &MongoStorage{
		Client:   client,
		Database: db,
	}
	return mongoStorage
}
