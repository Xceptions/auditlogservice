package database

import (
	"context"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// we have two db's we need to initialize. A postgres db
// and a mongodb. The postgres is used for storing user credentials
// but the mongodb is used for storing events

func ConnectPostgresDB() *gorm.DB {
	err := godotenv.Load(".env")
	dbURL := os.Getenv("POSTGRES_CONN")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	helpers.HandleErr(err)

	db.AutoMigrate(&models.User{})

	return db
}

func ConnectMongoDB() *mongo.Database {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	// Getting the DB Name from .env
	audit_db_name := os.Getenv("MONGO_DBNAME")

	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://kene:kenepass@127.0.0.1:27017")

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
	db := client.Database(audit_db_name)
	return db
}
