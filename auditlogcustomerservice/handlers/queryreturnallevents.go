package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/joho/godotenv"

	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/models"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/database"
)

// Receives: response and request writers
// returns all events in the database.
// Not advisable to run since
func (h handler) QueryReturnAllEvents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := godotenv.Load(".env")
	helpers.HandleErr(err)

    collectionName := os.Getenv("COLLECTION")

	userToken := r.Header.Get("x-access-token")

	var user models.User

	// will throw an error if user is not authorized
	helpers.IsAuthorized(w, userToken)

	DB := database.ConnectMongoDB()
	collection := DB.Collection(collectionName)

	findUser := h.DB.Where("Username = ?", userName).First(&user)
	if errors.Is(findUser.Error, gorm.ErrRecordNotFound) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}

	getUserAccount := h.DB.Where("user_id = ?", user.ID).First(&account)
	if errors.Is(getUserAccount.Error, gorm.ErrRecordNotFound) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Account not found")
	}

	// Send a 201 retrieved response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
