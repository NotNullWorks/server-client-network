package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message := "Hello from client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data to server:", err)
		return
	}

	fmt.Println("Data sent to server")

	// Read response from server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response from server:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("Response from server:", response)
}
