package main

import (
	udpPing "backend/udpPing"
	utils "backend/utils"
	"fmt"
	"net"
)

func main() {
	udpPing.InitUDPPingServer(9000)
	initTCPServer(9001)
}

func initTCPServer(port int) {
	listener, err := net.Listen("tcp", utils.ByteArraytoString(utils.GetLocalIP())+fmt.Sprintf(":%d",port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port: ",port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("connection from: "+conn.LocalAddr().String())
	}
}