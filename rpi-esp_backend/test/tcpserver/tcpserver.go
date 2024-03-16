package main

import (
	"fmt"
	"net"
)

func main() {
	i := 0
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		i++
		go handleClient(conn, i)
	}
}

func handleClient(conn net.Conn, i int) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			//   fmt.Println("Error:", err)
			break
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s from %d\n", buffer[:n], i)
	}
	fmt.Println("closed ", i)
}
