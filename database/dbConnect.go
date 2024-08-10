package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	MongoDB := os.Getenv("MONGODB_URL")

	if MongoDB == "" {
		log.Fatal("MONGODB_URL environment variable not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(MongoDB).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal("Error in connecting DB: ", err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	return client
}

var Client *mongo.Client = DbInstance()

func DisconnectDB() {
	if err := Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println(Client)
	fmt.Println("Disconnected")

}

func OpenCollection(collectionName string) *mongo.Collection {
	collection := Client.Database("cluster0").Collection(collectionName)

	return collection
}
