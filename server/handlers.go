package main

import (
	"bufio"
	"fmt"
	"net"
)

/** Listens for messages from a client and forwards to the broadcast channel. */
func handleClient(conn net.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr())
	}()

	/** First get the client username and add new client to clients map */
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		mu.Lock()
		clients[conn] = scanner.Text()
		mu.Unlock()
	}

	/**
	 * Create a scanner that reads data from the connection conn. It
	 * continuously reads the next line from the connection and returns
	 * true if there's another line or message.
	 */
	for scanner.Scan() {
		mu.Lock()
		username := clients[conn]
		mu.Unlock()
		msg := fmt.Sprintf("[%s] %s", username, scanner.Text())
		broadcast <- msg
	}
}

/** Waits for messages and then broadcasts them to every client. */
func handleBroadcast() {
	for msg := range broadcast {
		mu.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mu.Unlock()
	}
}
