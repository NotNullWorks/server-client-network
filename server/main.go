package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Received from client:", message)

		// Process the received message (e.g., parse JSON, perform operations)

		// Send a response back to the client
		response := []byte("Response from server: " + message)
		_, err = conn.Write(response)
		if err != nil {
			fmt.Println("Error sending response to client:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}
