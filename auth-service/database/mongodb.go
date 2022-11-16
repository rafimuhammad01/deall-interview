package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Database {
	ctx := context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI(os.Getenv("MONGO_DB_URI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("error from mongodb: ", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("error from mongodb: ", err)
	}

	return client.Database(os.Getenv("MONGO_DB_NAME"))
}
