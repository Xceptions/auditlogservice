package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/database"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Receives: response and request writers
// returns all events in the database.
// Not advisable to run since
func (h handler) QueryEventsByFieldAndValue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := godotenv.Load(".env")
	helpers.HandleErr(err)

	field := mux.Vars(r)["field"]
	value := mux.Vars(r)["value"]

	collectionName := os.Getenv("COLLECTION")

	// will have to comment out token authentication
	// useful during actual deployment
	// userToken := r.Header.Get("x-access-token")
	// helpers.IsAuthorized(w, userToken)

	DB := database.ConnectMongoDB()
	collection := DB.Collection(collectionName)

	filter := bson.D{{field, value}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var events []models.Event
	if err = cursor.All(context.TODO(), &events); err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"result": events})
}
