package main

//import packages
import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//NOTE: Server must be shutdown manually
	// host port number
	const address string = ":8080"

	// net.Listen return a Listener interface, and an error
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	//defer the closing of connection until main() completes
	defer listener.Close()

	//infinite loop, will keep server running, accepting connection, printing out message, the closing client connection
	for {
		//Accept connection for request at port 8080
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		//Connection actions/messages
		io.WriteString(connection, "\nYou've connected to the server at localhost" + address + "\n")
		fmt.Fprintln(connection, "Connection open.........")
		fmt.Fprintln(connection, "Closing connection......")

		//close client connection
		connection.Close()
	}

}
