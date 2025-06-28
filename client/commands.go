package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func handlClientSideCommand(input string, conn net.Conn) bool {
	/** Split input into a parts by " " and set extract the command */
	parts := strings.Fields(input)
	command := parts[0]

	switch command {
	case "/help":
		printHelp()
	case "/quit":
		fmt.Println("Goodbye!")
		conn.Close()
		os.Exit(0)
	default: // Forward server-related commands
		fmt.Fprintln(conn, input)
		return true
	}

	return true
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  /help                 Show this message")
	fmt.Println("  /quit                 Exit the chat")
	fmt.Println("  /name <username>      Change your nickname")
	fmt.Println("  /msg @user <message>  Send a private message")
}
