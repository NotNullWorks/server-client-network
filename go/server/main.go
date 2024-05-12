package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve UDP address and port
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server started. Waiting for messages...")

	// Buffer for incoming messages
	buffer := make([]byte, 1024)

	for {
		// Read from UDP connection
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			return
		}

		fmt.Printf("Received message from %s: %s\n", addr.String(), string(buffer[:n]))

		// Respond to the client
		response := []byte("Hello from server!")
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("Error sending response to client:", err)
			return
		}
	}
}
