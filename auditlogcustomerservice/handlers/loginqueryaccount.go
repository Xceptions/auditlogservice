package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Receives: response and request writers
// function receives input from the user containing username
// and password. We then use authenticate these values by
// first checking if the username exists in the database,
// then we compare the passwords. If we find a match, we
// generate a JWT, and return it
// Returns: string
func (h handler) LoginQueryAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)

	var userInput models.User
	var userDetail models.User
	err = json.Unmarshal(body, &userInput)
	helpers.HandleErr(err)

	findUser := h.DB.Where("Username = ?", userInput.Username).First(&userDetail)
	if errors.Is(findUser.Error, gorm.ErrRecordNotFound) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}

	// verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(userInput.Password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode("Wrong password")
	}

	tokenString := helpers.GenerateJWT(userInput.Username)

	// Send a 201 response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenString)
}
