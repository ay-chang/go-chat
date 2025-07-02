package main

import (
	"bufio"
	"fmt"
	"net"
)

/** Listens for messages from a single client and forwards to the broadcast channel. */
func handleClient(conn net.Conn) {
	var username string

	/** Handle exiting program for current client */
	defer func() {
		mu.Lock()
		delete(clients, conn)
		delete(users, username)
		mu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr())
	}()

	/**
	* First create a scanner and get the client username and add new client username to
	* clients map and also add the client to the user map.
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
	 * Continuously scan the next line from the connection and returns true if there's
	 * another line or message. Also checks for certain commands such as /msg, otherwise
	 * broadcast the message normally.
	 */
	for scanner.Scan() {
		mu.Lock()
		username := clients[conn]
		mu.Unlock()

		msg := scanner.Text()
		if handleServerSideCommands(msg, conn) {
			continue
		}
		broadcast <- fmt.Sprintf("[%s] %s", username, msg) // normal broadcast
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
