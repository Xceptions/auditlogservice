package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/database"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/handlers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
)

// returns 404 error for non-existent urls
func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// HTTP server will be used for querying
// audit logs. We receive the connection from the
// client, then initialize a new connection to the
// database. We also create versioned routes for the
// APIs. Rate-limiting and caching is applied to the
// querying handler

// TODO: Perform rate-limiting by client ip address
// because currently, I am rate-limiting the whole
// application
func spinUpHTTPServer() {
	log.Println("Starting HTTP Server...")

	DB := database.ConnectPostgresDB()
	h := handlers.New(DB)
	router := mux.NewRouter()

	var api = router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(error404)

	var apiVersion1 = api.PathPrefix("/v1").Subrouter()
	apiVersion1.NotFoundHandler = http.HandlerFunc(error404)

	apiVersion1.HandleFunc("/createuser", h.CreateQueryAccount).Methods(http.MethodPost)
	apiVersion1.HandleFunc("/loginuser", h.LoginQueryAccount).Methods(http.MethodPost)
	apiVersion1.HandleFunc("/getevents/{field}/{value}", helpers.RateLimiter(h.QueryEventsByFieldAndValue)).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

// starts servers
func main() {
	spinUpHTTPServer()
}
