package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	clients   = make(map[net.Conn]string) // all connected clients
	users     = make(map[string]net.Conn) // reverse hashmap of clients
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
