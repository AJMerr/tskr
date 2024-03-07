package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDotEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("MONGODB_URI")
}

var Client *mongo.Client

func DBConnect(uri string, database string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(GetDotEnv()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	return err
}

func DisconnectDB() error {
	return Client.Disconnect(context.Background())
}
