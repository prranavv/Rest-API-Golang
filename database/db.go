package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func StartMongodb() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return nil
}

func CloseMongodb() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return MongoClient.Database("test").Collection(name)
}
