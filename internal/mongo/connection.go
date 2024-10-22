package mongo

import (
	"context"
	"log"
	"os"
	"webhook-consumer/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var client *mongo.Client

const database = "webhook"

func init() {
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	var err error

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	logger.Logger.Info("Connected to MongoDB")

}
