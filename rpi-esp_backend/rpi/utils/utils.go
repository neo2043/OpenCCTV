package utils

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() []byte {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP
}

func ByteArraytoString(arr []byte) string {
	var aggregateString string
	for i, data := range arr {
		aggregateString += fmt.Sprintf("%d",int(data))
		if(i!=3){
			aggregateString+="."
		}
	}
	return aggregateString
}

func ReadfromClient(client *net.UDPConn) (string, string) {
	var buf [512]byte
	n, _, err := client.ReadFromUDP(buf[0:])
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(buf[:n]), ":") {
		str := strings.Split(string(buf[:n]), ":")
		return str[0], str[1]
	} else {
		return "", ""
	}
}