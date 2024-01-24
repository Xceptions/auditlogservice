package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"github.com/xceptions/golangauditlog/handlers"
	"github.com/gorilla/mux"
)



// TCP server will be used for accepting
// audit log
func spinUpTCPServer() {
	const 
	(
	CONN_HOST = "localhost"
	CONN_PORT = "8002"
	CONN_TYPE = "tcp"
	)

	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error starting tcp server : ", err)
	}
	defer listener.Close()
	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for 
	{
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go handlers.HandleLog(conn)
	}

}

// HTTP server will be used for querying
// audit logs
func spinUpHTTPServer() {
	log.Println("Starting HTTP Server...")

    DB := database.ConnectDB()
	h := handlers.New(DB)
    router := mux.NewRouter()

	router.HandleFunc("/createuser", h.CreateUserAccount).Methods(http.MethodPost)
	router.HandleFunc("/login", h.LogInUser).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

// starts two servers
func main() {
	
}