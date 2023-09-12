package common

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func InitMongoDB() {
	// Set the MongoDB URI
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	uri := os.Getenv("MONGO_URI")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a MongoDB client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("levels")

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("failed to connect to MongoDB!", err)
	}

	fmt.Println("Connected to MongoDB!")
}

func CleanupMongoDB(client *mongo.Client, ctx context.Context) {
	client.Disconnect(ctx)
}

func GetDBCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
