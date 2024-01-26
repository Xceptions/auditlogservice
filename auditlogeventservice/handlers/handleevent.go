package handlers

import (
	"fmt"
	"net"
)

func HandleEvent(bufferedChannel chan []byte, conn net.Conn) {
	defer conn.Close()

	go addEventToBuffer(bufferedChannel, conn)

	for msg := range bufferedChannel {
		fmt.Println("the message is:", string(msg))
	}
}
