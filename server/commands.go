package main

import (
	"fmt"
	"net"
	"strings"
)

func handleServerSideCommands(msg string, conn net.Conn, username string) bool {
	parts := strings.Fields(msg)
	command := parts[0]

	switch command {
	case "/msg":
		recipient := strings.TrimPrefix(parts[1], "@")
		privateMsg := strings.Join(parts[2:], " ")

		mu.Lock()
		targetConn, ok := users[recipient]
		mu.Unlock()

		if ok {
			sender := clients[conn]
			fmt.Fprintf(targetConn, "[private] %s: %s\n", sender, privateMsg) // send private msg
			fmt.Fprintf(conn, "[you → @%s] %s\n", recipient, privateMsg)      // echo back private msg
		} else {
			fmt.Fprintf(conn, "User %s not found.\n", recipient)
		}
	default: // Forward server-related commands
		fmt.Printf("Forwarding message: %s\n", msg)
		fmt.Fprintln(conn, msg)
		return false
	}

	return true
}
