package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve UDP address and port
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	// Send message to server
	message := []byte("Hello from client!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending message to server:", err)
		return
	}

	fmt.Println("Message sent to server.")

	// Receive response from server
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading response from server:", err)
		return
	}

	fmt.Println("Response from server:", string(buffer[:n]))
}
