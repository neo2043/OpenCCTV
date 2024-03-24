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
    c:=0
	for{
        data := []byte("Hello, Server!"+fmt.Sprintf("%d",c))
		_, err = conn.Write(data)
		if err != nil {
		 fmt.Println("Error:", err)
		 return
		}
        time.Sleep(time.Second)
        c++
	}
    // Read and process data from the server
    // ...
}