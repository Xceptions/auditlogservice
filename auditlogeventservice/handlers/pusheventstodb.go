package handlers

import (
	"context"
	"log"
	"os"

	"encoding/json"
	"time"

	"github.com/joho/godotenv"
	"github.com/xceptions/auditlogservice/auditlogeventservice/database"
	"github.com/xceptions/auditlogservice/auditlogeventservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogeventservice/models"
)

// receives events in groups of insertManyLimit (e.g. 5)
// prepares them as an insertmany statement,
// adds them to the mongodb database
// the eventJson var is used to unmarshal each event to
// the appropriate json value
// the eventToPush slice is used to hold all the events.
// This is what is eventually pushed to the db
func PushEventToDB(eventsSlice [][]byte) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	DB := database.ConnectMongoDB()
	collection := DB.Collection(os.Getenv("COLLECTION"))
	var eventsToPush []interface{}

	for _, event := range eventsSlice {
		eventObject := models.Event{}
		err := json.Unmarshal(event, &eventObject)
		helpers.HandleErr(err)
		eventsToPush = append(eventsToPush, models.Event{Customer: eventObject.Customer, EventType: eventObject.EventType, Time: time.Now(), Specifics: eventObject.Specifics})
	}

	_, err = collection.InsertMany(context.TODO(), eventsToPush)
	if err != nil {
		panic(err)
	}
}
