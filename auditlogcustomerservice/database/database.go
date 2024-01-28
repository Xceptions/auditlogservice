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

// var db *mongo.Database

func ConnectPostgresDB() *gorm.DB {
	dbURL := "postgresql://kene:kenepass@localhost:5432/postgres"

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
	audit_db_name := os.Getenv("AUDIT_DB_NAME")

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
