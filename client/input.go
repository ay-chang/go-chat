package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/**
 * Wait for input: Wraps keyboard input in a scanner so we can read the
 * keyboard inputs (os.Stdin) one line at a time. Then continuously
 * run (scan) for messages to send to our TCP connection conn.
 */
func handleUserInput(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ") // show the inital input prompt
		scanner.Scan()
		text := scanner.Text()

		if handleCommand(text) {
			continue
		} else {
			fmt.Fprintln(conn, text)
		}

		/** Clear the line that the user just typed */
		fmt.Print("\033[1A")
		fmt.Print("\033[2K\r")
	}
}
