package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/xceptions/golangauditlog/database"
	"github.com/xceptions/golangauditlog/helpers"
	"github.com/xceptions/golangauditlog/models"
)

// Receives: response and request writers
// function receives input from the user containing username, email
// and password. We then use these values to populate both the
// user table and account table. The account is generated from the
// user detail. It returns the status of the account creation
// Returns: string
func CreateQueryAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DB := database.ConnectDB()

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)

	var user models.User
	json.Unmarshal(body, &user)

	generatePassword := helpers.HashAndSalt([]byte(user.Password))
	user.Password = generatePassword
	result, err := DB.Collection("QueryUsers").InsertOne(ctx, &user)
	helpers.HandleErr(err)

	// Send a 201 created response
	tokenString := helpers.GenerateJWT(user.Username)

	// Send a 201 response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "token": tokenString})
}
