package main

import (
	// "fmt"
	"crypto/rand"
	"log"
	"net"

	"github.com/xceptions/auditlogservice/auditlogeventservice/handlers"
)

// TCP server will be used for accepting
// audit log
func spinUpTCPServer() {
	log.Println("spinning up tcp server")
	token := make([]byte, 4)
	rand.Read(token)
	bufferedChannel := make(chan []byte, 4)
	bufferedChannel <- token
	bufferedChannel <- token
	bufferedChannel <- token

	const (
		CONN_HOST = "localhost"
		CONN_PORT = "8952"
		CONN_TYPE = "tcp"
	)

	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error starting tcp server : ", err)
	}
	defer listener.Close()
	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go handlers.HandleEvent(bufferedChannel, conn)
	}
}

// starts servers
func main() {
	spinUpTCPServer()
}
