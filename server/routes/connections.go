package routes

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

// DBInstance func
func DBInstance() *mongo.Client {
	// load the .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to mongoDB client
	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

// Client database instance
var Client *mongo.Client = DBInstance()

// OpenCollection Open Collection is a function that makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	return collection
}
