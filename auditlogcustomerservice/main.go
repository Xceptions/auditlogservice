package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/database"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/handlers"
)

// returns 404 error for non-existent urls
func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// HTTP server will be used for querying
// audit logs
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
	apiVersion1.HandleFunc("/getevents/{field}/{value}", h.QueryEventsByFieldAndValue).Methods(http.MethodGet)

	// router.HandleFunc("/createuser", h.CreateQueryAccount).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

// starts servers
func main() {
	spinUpHTTPServer()
}
