// package tcphttpServer

// import (
// 	"fmt"
// 	// "github.com/gofiber/fiber/v2"
// 	"net"
// )

// func initHTTPServer(port int) {
// 	// app := fiber.New()
// 	// app.Get("/update",func(c *fiber.Ctx) error {
// 	// 	c.SendFile()
// 	// })
// }

// func GetLocalAddr() *net.UDPAddr {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer conn.Close()
// 	conn.LocalAddr().String()
// }

// func initTCPServer(port string) {
// 	listener, err := net.Listen("tcp", "localhost:"+port)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server is listening on port " + port)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			continue
// 		}
// 		go handleClient(conn)
// 	}
// }

// func handleClient(conn net.Conn) {
// 	defer conn.Close()

// 	buffer := make([]byte, 1024)

// 	for {
// 		n, err := conn.Read(buffer)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}

// 		fmt.Printf("Received: %s\n", buffer[:n])
// 	}
// }

package main

import (
	fio "backend/fileio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net"
)

type serverObjectStruct struct {
	tcpListener  net.Listener
	httpListener *fiber.App	
}

func (t *serverObjectStruct) setTCPServer(address,port string) {
	var err error
	t.tcpListener, err = net.Listen("tcp", address+port)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *serverObjectStruct) setHTTPServer() {
	t.httpListener = fiber.New()
}

