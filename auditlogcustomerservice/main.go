package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/handlers"
	"github.com/xceptions/auditlogservice/auditlogcustomerservice/database"
)

// HTTP server will be used for querying
// audit logs
func spinUpHTTPServer() {
	log.Println("Starting HTTP Server...")

	DB := database.ConnectPostgresDB()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/createuser", h.CreateQueryAccount).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

// starts servers
func main() {
	spinUpHTTPServer()
}
