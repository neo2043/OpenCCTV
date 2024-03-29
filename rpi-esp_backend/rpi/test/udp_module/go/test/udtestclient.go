package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func getMyAddr() *net.UDPAddr {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.LocalAddr().(*net.UDPAddr).Port=9000
	return conn.LocalAddr().(*net.UDPAddr)
}

func main(){
	serverAddr, err := net.ResolveUDPAddr("udp4", "192.168.1.255:9000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count:=0
	for{
		fmt.Println("udpclient: "+fmt.Sprintf("%d",count))
		// _, err = conn.Write([]byte("udpclient2 "+fmt.Sprintf("%d",count)))
		_, err = conn.Write([]byte(fmt.Sprintf("%d",count)))
		count++
		time.Sleep(time.Millisecond*500)
	}

}