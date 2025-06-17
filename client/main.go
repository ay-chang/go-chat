package main

import (
	"net"
)

var (
	username string = ""
)

func main() {
	/**
	 * Creates a client side TCP connection to the server localhost:9000
	 * and if the connection is succesful, we get a conn object that we
	 * can use to write messages to the server
	 */
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ShowWelcomeMessage()
	createUser(conn)

	/** Receive messages while being able to send out messages */
	go receiveMessages(conn)

	/** Continuously wait for user input*/
	handleUserInput(conn)
}
