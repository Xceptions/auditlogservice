package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectMongoDB() *mongo.Database {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	// Getting the DB Name from .env
	mongo_db := os.Getenv("MONGO_DBNAME")

	// set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// set the database and collection variables
	db = client.Database(mongo_db)
	return db
}
