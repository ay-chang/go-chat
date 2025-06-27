package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

/** Listens for messages from a client and forwards to the broadcast channel. */
func handleClient(conn net.Conn) {
	var username string

	defer func() {
		mu.Lock()
		delete(clients, conn)
		delete(users, username)
		mu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr())
	}()

	/**
	 * First get the client username and add new client username to clients map and also
	 * add the client to the user map
	 */
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		username = scanner.Text()

		mu.Lock()
		clients[conn] = username
		users[username] = conn
		mu.Unlock()
	}

	/**
	 * Create a scanner that reads data from the connection conn. It continuously reads
	 * the next line from the connection and returns true if there's another line or
	 * message. Also checks for certain commands such as /msg, otherwise broadcast the
	 * message normally.
	 */
	for scanner.Scan() {
		mu.Lock()
		username := clients[conn]
		mu.Unlock()

		msg := scanner.Text()
		if strings.HasPrefix(msg, "/msg ") {
			parts := strings.Fields(msg)
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

		} else {
			broadcast <- fmt.Sprintf("[%s] %s", username, msg) // normal broadcast
		}
	}
}

/** Waits for messages from handleClient and then broadcasts them to every client. */
func handleBroadcast() {
	for msg := range broadcast {
		mu.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mu.Unlock()
	}
}
