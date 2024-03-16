package main

import (
	"fmt"
	"backend/udpPing"
)

func main() {
	// udpPing.InitServer(9000)
	f:=udpPing.UdpAddr{}
	f.GetLocalAddr()
	f.SetPort(8000)
	fmt.Println(f.GetUDP_IP())
	
}