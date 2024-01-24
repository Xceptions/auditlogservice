package handlers

import (
	"bufio"
	"fmt"
	"net"
)

func HandleLog(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message Received from the client: ", string(message))
	conn.Close()
}
