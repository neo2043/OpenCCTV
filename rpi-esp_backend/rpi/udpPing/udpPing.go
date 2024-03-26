package udpPing

import (
	udpPingStructs "backend/structs/udppingstruct"
	"backend/utils"
	"fmt"
	"net"
)

func InitUDPPingServer(port int) {
	addr := udpPingStructs.UdpAddr{}
	addr.SetIP(utils.GetLocalIP())
	addr.SetPort(port)
	clientRecieverConn, err := net.ListenUDP("udp4", addr.GetUDPAddrStruct())
	fmt.Println("UDP server started on port: ",port)
	if err != nil {
		fmt.Println("server reciever error")
	}
	go func() {
		for {
			clientAddr, encData := utils.ReadfromClient(clientRecieverConn)
			if encData == "INITCONN" {
				clientUDPAddr, err := net.ResolveUDPAddr("udp4", clientAddr)
				if err != nil {
					fmt.Println(err)
				}
				conn, err := net.DialUDP("udp4", nil, clientUDPAddr)
				if err != nil {
					fmt.Println(err)
				}
				conn.Write([]byte(addr.GetIP_Port() + ":connect"))
				conn.Close()
			}
		}
	}()
}