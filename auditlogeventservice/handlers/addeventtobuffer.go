package handlers

import (
	// "bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func addEventToBuffer(bufferedChannel chan []byte, conn net.Conn) {

	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read error: ", err)
			continue
		}
		if err == io.EOF {
			// We reached EOF, let’s close that goroutine
			log.Println("reading error:", err)
			break
		}
		if errors.Is(err, io.EOF) { // prefered way by GoLang doc
			fmt.Println("Reading file finished...")
			break
		}

		// fmt.Println(string(buffer[:n]))

		bufferedChannel <- buffer[:n]
	}
}
