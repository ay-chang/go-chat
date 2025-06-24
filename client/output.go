package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/**
 * Continuously listen for incoming messages from the server connection
 * and print each message to the terminal as it's received.
 */
func receiveMessages(conn net.Conn) {
	/** Prepare terminal ui for next message and then receive messages */
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Print("\r")
		fmt.Println(scanner.Text())
		fmt.Print("> ")
	}

	/**
	 * When the server closes, this function will see that conn
	 * has been closed so scanner.Scan returns false and the loops ends,
	 * so exit the program
	 */
	fmt.Println("Disconnected from server.")
	os.Exit(0) // stop the whole program
}

/** Receive private messages */
