package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var Username string

/** Allows the client to give itself a username */
func createUser(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your username: ")

	if scanner.Scan() {
		username = scanner.Text()
		fmt.Fprintln(conn, username)
	}
}
