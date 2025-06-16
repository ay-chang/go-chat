# 🧵 Go-Terminal-Chat

A simple terminal-based chat application written in Go — built as a learning project to explore **network programming**, **concurrency**, and **message broadcasting** using core Go features.

This is a basic but functional chat app that allows multiple clients to connect to a server, send messages, and receive messages from others in real time. Over time, I plan to add new features like nicknames, private messaging, command parsing, and more.

## ✨ Features (Current)

- Multi-client TCP server using the `net` package
- Concurrent client handling with goroutines
- Real-time message broadcasting using Go channels
- Basic terminal client that can send and receive messages
- Graceful client disconnection handling

Foundational coding techniques used:

### ⚙️ TCP Networking in Go

- Using `net.Listen()` and `net.Dial()` to set up client-server communication
- Managing TCP connections as byte streams

### 🧵 Goroutines and Concurrency

- Handling multiple clients at once without blocking
- Running background listeners (e.g. receive messages while typing)

### 🔁 Channels and Message Broadcasting

- Centralizing messages from all clients into a broadcast loop
- Distributing messages out to all connected clients safely

### 🔐 Safe Shared State

- Using `sync.Mutex` to prevent race conditions when accessing shared resources (like the clients map)

## 🚀 Planned Features

- Nicknames and user identification
- Command support (`/help`, `/msg @user`, `/quit`)
- Colored terminal output per user
- Chat history logging
- Hosting on a public server

## 🛠️ How to Run Locally

### Server

```bash
cd server
go run main.go
```

### Client (in another terminal)

```bash
cd client
go run main.go
```
