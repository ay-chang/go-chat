package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients   = make(map[net.Conn]string) // all connected clients
	broadcast = make(chan string)         // channel to broadcast messages
	mu        sync.Mutex                  // protects the clients map
)

func main() {
	/** Start a tcp server listening on port 9000 */
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 9000...")

	go handleBroadcast()

	for {
		/**
		 * Wait for client to connect, blocking so code stops here until
		 * a connection is made. When a connection is made, returns a conn
		 * which is an active connection to a client.
		 */
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		/** Add new client */
		mu.Lock()
		clients[conn] = conn.RemoteAddr().String()
		mu.Unlock()

		fmt.Println("New client connected: ", conn.RemoteAddr())
		go handleClient(conn)
	}
}

/** Listens for messages from a client and forwards to the broadcast channel. */
func handleClient(conn net.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr())
	}()

	/**
	 * Create a scanner that reads data from the connection conn. It
	 * continuously reads the next line from the connection and returns
	 * true if there's another line or message.
	 */
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := fmt.Sprintf("[%s] %s", conn.RemoteAddr(), scanner.Text())
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
