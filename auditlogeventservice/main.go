package main

import (
	"fmt"
	"log"
	"net"

	"github.com/xceptions/auditlogservice/auditlogeventservice/handlers"
	"github.com/xceptions/auditlogservice/auditlogeventservice/helpers"
)

// TCP server will be used for accepting
// audit log
func spinUpTCPServer(bufferedChannel chan []byte) {
	log.Println("spinning up tcp server...")

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

		if helpers.TCPConnectionIsAuthenticated(conn) {
			go handlers.PushEventsToBuffer(bufferedChannel, conn)
		}
	}
}

// starts servers, creates a buffered channel that is threadsafe
// spins up a tcp server, then watches the buffered channel
// for when we have our insert limit
func main() {
	// general channel that module will make use of
	bufferedChannel := make(chan []byte, 4)
	go spinUpTCPServer(bufferedChannel)

	// deciding limit to use for insertMany operation
	insertManyLimit := 5

	// holding the events in bulk, a second buffer
	// will go ahead to initialize it here
	var eventsSlice [][]byte
	eventsSlice = [][]byte{}

	for events := range bufferedChannel {
		eventsSlice = append(eventsSlice, events)
		if len(eventsSlice) == insertManyLimit {
			go handlers.PushEventToDB(eventsSlice)
			eventsSlice = [][]byte{} // clear the event slice
			fmt.Println("the len of eventsSize in the main thread is: ", len(eventsSlice))
		}
	}
}
