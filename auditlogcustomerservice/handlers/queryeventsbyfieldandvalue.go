package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/database"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Receives: response and request writers
// returns all events in the database where
// a certain field has the value of a certain
// value
func (h handler) QueryEventsByFieldAndValue(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	err := godotenv.Load(".env")
	helpers.HandleErr(err)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // Redis server address
		Password: "",                      // No password
		DB:       0,                       // Default DB
	})

	val, err := redisClient.Get(context.Background(), r.URL.Path).Result()

	if err == nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(val)
	}

	field := mux.Vars(r)["field"]
	value := mux.Vars(r)["value"]

	collectionName := os.Getenv("COLLECTION")

	// will have to comment out token authentication
	// useful during actual deployment
	// userToken := r.Header.Get("x-access-token")

	// helpers.IsAuthorized(w, userToken)

	DB := database.ConnectMongoDB()
	collection := DB.Collection(collectionName)

	filter := bson.D{primitive.E{Key: field, Value: value}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var events []models.Event
	if err = cursor.All(context.TODO(), &events); err != nil {
		panic(err)
	}
	result := map[string]interface{}{"result": events}
	resultMarshalled, _ := json.Marshal(result)

	err = redisClient.Set(context.Background(), r.URL.Path, resultMarshalled, 1*time.Hour).Err()
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultMarshalled)
}
