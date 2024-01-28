package handlers

import (
	"fmt"
	"net"
)

// receives the channel and conn from main method
// creates a buffer to hold the data from the conn
// parses it up to a defined byte
// adds the parsed data to the bufferedChannel.
// this buffer acts as a temp store for the data
func PushEventsToBuffer(bufferedChannel chan []byte, conn net.Conn) {
	defer conn.Close()
	fmt.Println("pushing messages into buffer...")

	buffer := make([]byte, 2048) // 2048 = assumption for reasonable event byte size

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read error: ", err)
			continue // to continue listening for connections
		}
		bufferedChannel <- buffer[:n]
	}
}
