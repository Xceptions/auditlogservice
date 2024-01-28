package handlers

import (
	"context"
	"log"

	"encoding/json"
	"fmt"
	"time"

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
	fmt.Println("The len of the event slice passed into the pushevent to db method is: ", len(eventsSlice))

	DB := database.ConnectDB()
	collection := DB.Collection("Events")

	var eventsToPush []interface{}

	for _, event := range eventsSlice {
		fmt.Println(event)
		fmt.Println(string(event))
		// eventJson, err := json.Marshal(event)
		eventObject := models.Event{}
		err := json.Unmarshal(event, &eventObject)
		if err != nil {
			log.Printf("error decoding sakura response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("sakura response: %q", eventObject)
		}
		fmt.Println(eventObject)
		helpers.HandleErr(err)
		eventsToPush = append(eventsToPush, models.Event{Customer: eventObject.Customer, EventType: eventObject.EventType, Time: time.Now(), Specifics: eventObject.Specifics})
	}

	_, err := collection.InsertMany(context.TODO(), eventsToPush)
	if err != nil {
		panic(err)
	}
}
