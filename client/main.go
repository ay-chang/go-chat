package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	/**
	 * Wait for input: Wraps keyboard input in a scanner so we can read the
	 * keyboard inputs (os.Stdin) one line at a time. Then continuously
	 * run (scan) for messages to send to our TCP connection conn.
	 */
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ") // show the inital input prompt
		scanner.Scan()
		fmt.Fprintln(conn, scanner.Text())

		// Clear the line you just typed
		fmt.Print("\033[1A")
		fmt.Print("\033[2K\r")
	}

}

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

/** Allows the client to give itself a username */
func createUser(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your username: ")

	if scanner.Scan() {
		username = scanner.Text()
		fmt.Fprintln(conn, username)
	}
}

func ShowWelcomeMessage() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      👋 Welcome to Go Chat!                ║")
	fmt.Println("║       A lightweight terminal-based chat experience.        ║")
	fmt.Println("║             Type your messages and join the convo.         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}
