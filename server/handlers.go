package main

import (
	"bufio"
	"fmt"
	"net"
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
		if handleServerSideCommands(msg, conn, username) {
			continue
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
