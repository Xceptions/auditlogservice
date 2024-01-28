package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/models"
)

// Receives: response and request writers
// function receives input from the user containing username, email
// and password. We then use these values to populate both the
// user table and account table. The account is generated from the
// user detail. It returns the status of the account creation
// Returns: string
func (h handler) CreateQueryAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)

	var user models.User
	json.Unmarshal(body, &user)

	generatePassword := helpers.HashAndSalt([]byte(user.Password))
	user.Password = generatePassword
	h.DB.Create(&user)

	// Send a 201 created response
	tokenString := helpers.GenerateJWT(user.Username)
	fmt.Println(tokenString)

	// Send a 201 response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenString)

	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "token": tokenString})
}
