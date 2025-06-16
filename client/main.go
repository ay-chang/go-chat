package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	fmt.Println("Connected to server. Type your messages:")

	/** Receive messages while being able to send out messages */
	go receiveMessages(conn)

	/**
	 * Wraps keyboard input in a scanner so we can read the keyboard
	 * inputs (os.Stdin) one line at a time. Then continuously run (scan)
	 * for messages to send to our TCP connection conn.
	 */
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
}

/**
 * Continuously listen for incoming messages from the server connection
 * and print each message to the terminal as it's received.
 */
func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	/**
	 * When the server closes, this function will see that conn
	 * has been closed so scanner.Scan returns false and the loops ends,
	 * so exit the program
	 */
	fmt.Println("Disconnected from server.")
	os.Exit(0) // stop the whole program
}
