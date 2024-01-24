package main

import (
	// "fmt"
	"log"
	// "net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xceptions/golangauditlog/handlers"
)

// TCP server will be used for accepting
// audit log
// func spinUpTCPServer() {
// 	fmt.Println("spinning up tcp server")
// 	const
// 	(
// 	CONN_HOST = "localhost"
// 	CONN_PORT = "8002"
// 	CONN_TYPE = "tcp"
// 	)

// 	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
// 	if err != nil {
// 		log.Fatal("Error starting tcp server : ", err)
// 	}
// 	defer listener.Close()
// 	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Fatal("Error accepting: ", err.Error())
// 		}
// 		go handlers.HandleLog(conn)
// 	}
// }

// HTTP server will be used for querying
// audit logs
func spinUpHTTPServer() {
	log.Println("Starting HTTP Server...")

	// DB := database.ConnectDB()
	// h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/createuser", handlers.CreateQueryAccount).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

// starts two servers
func main() {
	spinUpHTTPServer()
}
