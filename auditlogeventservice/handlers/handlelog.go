package handlers

import (
	// "bufio"
	"fmt"
	"io"
	"log"
	"net"
	// "strconv"
)

// func saveLog(customer string, event string, specifics interface{}) {
// 	input := models.AuditModel{
// 		Customer:  customer,
// 		Event:     event,
// 		Time:      time.Now(),
// 		Specifics: specifics,
// 	}

// 	logEntry, _ := json.Marshal(input)
// 	file, _ := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	defer file.Close()

// 	file.WriteString(string(logEntry) + "\n")
// }

// func writeLogToDB() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	DB := database.ConnectDB()
// 	file, err := os.Open("audit.log")
// 	if err != nil {
// 		log.Fatalf("Error opening audit.log: %s", err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var bulkWriteOperations []mongo.WriteModel
// 	for scanner.Scan() {
// 		var entry models.AuditModel
// 		if err := json.Unmarshal([]byte(scanner.Text()), &entry); err != nil {
// 			log.Fatalf("Error parsing JSON: %s", err)
// 		}

// 		body, err := json.Marshal(entry)
// 		if err != nil {
// 			log.Fatalf("Error marshaling entry to JSON: %s", err)
// 		}

// 		toWrite := mongo.NewInsertOneModel().SetDocument(models.AuditModel{Customer: body.Customer, Event: body.Event, Time: body.Time, Specifics: body.Specifics})
// 		bulkWriteOperations = append(bulkWriteOperations, toWrite)
// 	}
// 	opts := options.BulkWrite().SetOrdered(false)

// 	results, err := DB.Collection("WordToId").BulkWrite(ctx, bulkWriteOperations, opts)
// 	helpers.HandleErr(err)

// 	log.Println(results)

// }
// func HandleLog(conn net.Conn) {
// 	var bulkWriteOperations []mongo.WriteModel
// 	DB := database.ConnectDB()
// 	reader := bufio.NewReader(conn)

// 	for {
// 		event, err := reader.ReadString('\n')
// 		helpers.HandleErr(err)
// 		var body models.AuditModel
// 		json.Unmarshal([]byte(event), &body)
// 		toWrite := models.AuditModel{Customer: body.Customer, Event: body.Event, Time: body.Time, Specifics: body.Specifics}
// 		results, err := bufio.Writer(DB.Collection("AuditModel").InsertOne(ctx, toWrite))
// 		// fmt.Print("Message Received from the client: ", string(message))
// 	}

// 	helpers.HandleErr(err)
// 	log.Println(results)

// 	bufio.NewWriter(bulkWriteOperations)

// 	conn.Close()

// 	// input, err := io.ReadAll(conn)
// 	// helpers.HandleErr(err)
// 	// log.Println(string(input))

// 	// saveLog()

// }

type msgFormat struct {
	text []byte
	net.Conn
}

var accounts = make(map[net.Conn]int)
var conns = make(chan net.Conn)
var dconns = make(chan net.Conn)
var msgs = make(chan msgFormat)
var i int

func putDataInChannel(eventChannel chan []byte, dataConn net.Conn) {
	buffer := make([]byte, 1024)
	var dataLen int

	for {
		// Read data from the client
		n, err := dataConn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		dataLen = n
	}
	data := buffer[:dataLen]
	eventChannel <- data
	close(eventChannel)

}

// 	for {
// 		dataLen := make([]byte, 100)
// 		_, err := dataReader.Read(dataLen)
// 		if err != nil {
// 			break
// 		}
// 		intLen, _ := strconv.Atoi(string(dataLen))
// 		data := make([]byte, intLen)
// 		_, err = dataReader.Read(data)
// 		if err != nil {
// 			break
// 		}

// 	}
// 	close(eventChannel)
// 	// dconns <- conn
// }

func HandleLog(unbufferedChannel chan []byte, conn net.Conn) {
	defer conn.Close()
	fmt.Println("gets here")

	go reader(unbufferedChannel, conn)
	// for input := range unbufferedChannel {
	// 	fmt.Println("the input is ", string(input))
	// }
	fmt.Println("now here")
}

func reader(unbufferedChan chan []byte, conn net.Conn) {
	// We close the connection at any exit points
	defer conn.Close()

	// _, err := conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	// if err != nil {
	//     panic(err)
	// }

	// we read conn data in a 256 bytes slice
	b := make([]byte, 256)

	for {
		n, err := conn.Read(b)
		// fmt.Println(n)
		if err == io.EOF {
			// We reached EOF, let’s close that goroutine
			log.Println("reading error:", err)
			// break
		}
		if err != nil {
			log.Println("reading error:", err)
			// continue
		}
		// fmt.Printf("the string is %s", string(b[:n]))

		// We copy the current byte slice to a new byte slice that is passed through the mux channel
		out := make([]byte, n)
		copy(out, b)
		fmt.Printf("the string is %s", string(out))
		// We send the new byte slice back to the main goroutine that ranges over the muxChan data
		unbufferedChan <- out
	}
}
