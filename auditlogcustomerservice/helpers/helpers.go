package helpers

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Receives: err
// will be used to handle all of our errors and
// in order to obey the DRY principle
// Returns: None
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Receives: byte
// method is used to hash passwords to cryptic form
// in order to prevent password leak incase of a data
// breach
// Returns: string
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

// Receives: string (username)
// function is used to generate a JWT for a new user
// that logs in, populate it with the user's detail
// and send to the client.
// Returns: string
func GenerateJWT(Username string) string {
	var secretKey = []byte("kenechukwusecret")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = Username

	tokenString, err := token.SignedString(secretKey)
	HandleErr(err)

	return tokenString
}

// Receives: httResponseWriter, string
// function validates the users JWT and returns the username
// contained in the claims
// Returns: string
func IsAuthorized(w http.ResponseWriter, userToken string) string {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("kenechukwusecret"), nil
	})
	HandleErr(err)
	user := fmt.Sprintf("%v", claims["user"])

	return user
}

// rate limiting to help us limit the number of calls to our
// api
func RateLimiter(next http.HandlerFunc) http.HandlerFunc {
	limiter := rate.NewLimiter(10, 12)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		} else {
			next(w, r)
		}
	})
}
