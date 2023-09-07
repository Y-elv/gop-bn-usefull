package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Blog application!")
}

func main() {

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

	// Disconnect from the MongoDB server when the main function exits
	defer client.Disconnect(ctx)

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	// Now you can use the 'client' to interact with the MongoDB database

	// then time for using routers

	r := mux.NewRouter()

	r.HandleFunc("/", WelcomeHandler)
	port := "8000"
	fmt.Printf("Server running on port %s ...\n", port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
