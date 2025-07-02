package main

import (
	"fmt"
	"net"
	"strings"
)

/** Handle commands if detected, otherwise forward it to the broadcast channel */
func handleServerSideCommands(msg string, conn net.Conn) bool {
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
	case "/who":
		fmt.Fprintln(conn, "Here is a list of active users:")
		mu.Lock()
		for activeUser := range users {
			fmt.Fprintf(conn, "  - %s\n", activeUser)
		}
		mu.Unlock()
	default: // Forward server-related commands
		fmt.Fprintln(conn, msg)
		return false
	}

	return true
}
