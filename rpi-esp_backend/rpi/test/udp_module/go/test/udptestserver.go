package main

import (
	"fmt"
	"net"
	// "strings"
)

func getMyAddr() *net.UDPAddr {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		panic(err)
	}
	conn.LocalAddr().(*net.UDPAddr).Port=9000
	// conn.LocalAddr().(*net.UDPAddr).IP=net.ParseIP("100.100.146.110")
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr)
}

func main(){
	conn, err := net.ListenUDP("udp4", getMyAddr())
	if err != nil {
		panic(err)
	}
	println("== Listening on", getMyAddr().String())

	for {
		var buf [512]byte
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			panic(err)
		}

		// fmt.Println("-> Got message from ",sender.String()," message",string(buf[:n]))
		str:=string(buf[:n])
		fmt.Println(str)
	}
}