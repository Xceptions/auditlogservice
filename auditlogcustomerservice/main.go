package main

import (
	// "fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xceptions/golangauditlog/handlers"
)

// HTTP server will be used for querying
// audit logs
func spinUpHTTPServer() {
	log.Println("Starting HTTP Server...")

	router := mux.NewRouter()

	router.HandleFunc("/createuser", handlers.CreateQueryAccount).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

// starts two servers
func main() {
	spinUpHTTPServer()
}
