package main

import (
    "fmt"
    "net"
	"time"
)

func main() {
    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    // Send data to the server
    // ...
	data := []byte("Hello, Server!")
	for{
		_, err = conn.Write(data)
		if err != nil {
		 fmt.Println("Error:", err)
		 return
		}
		time.Sleep(time.Millisecond*1000)
	}
    // Read and process data from the server
    // ...
}