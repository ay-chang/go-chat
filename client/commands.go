package main

import (
	"fmt"
	"os"
	"strings"
)

func handleCommand(input string) bool {
	if !strings.HasPrefix(input, "/") {
		return false
	}

	/** Split input into a parts by " " and set extract the command */
	parts := strings.Fields(input)
	command := parts[0]

	switch command {
	case "/help":
		printHelp()
	case "/msg":
		messagePrivately(parts)
	default:
		fmt.Println("Unknown command:", command)
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

func messagePrivately(parts []string) {
	recipient := parts[1]
	message := strings.Join(parts[2:], " ")
	fmt.Fprintf(os.Stdout, "The recipient is %s and the message is %s\n", recipient, message)
}
