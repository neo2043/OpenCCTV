package udpPing

import (
	"fmt"
	"net"
	"strings"
)

type UdpAddr struct {
	addr *net.UDPAddr
}

func (t *UdpAddr) SetPort(port int) {
	t.addr.Port=port
}

func (t *UdpAddr) SetUDPIP(ip interface{}) {
	switch argType := ip.(type) {
		case string:
			if t.addr!=nil{
				t.addr.IP=net.ParseIP(ip.(string))
			}else{
				t.addr=&net.UDPAddr{}
				t.addr.IP=net.ParseIP(ip.(string))
			}
		case []byte:
			ip := ip.([]byte)
			if t.addr!=nil{
				t.addr.IP=net.IPv4(ip[0],ip[1],ip[2],ip[3])
			}else{
				t.addr=&net.UDPAddr{}
				t.addr.IP=net.IPv4(ip[0],ip[1],ip[2],ip[3])
			}
		default:
			fmt.Printf("Don't use type %T\n", argType)
	}
}

func (t *UdpAddr) GetIP() *net.UDPAddr {
	if t.addr==nil{
		fmt.Println("t.addr is nil")
	}
	return t.addr
}

func (t *UdpAddr) GetLocalUDPAddr() *net.UDPAddr {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	t.addr = conn.LocalAddr().(*net.UDPAddr)
	return t.addr
}

func InitUDPPingServer(port int) {
	addr:=&UdpAddr{}
	addr.GetLocalUDPAddr()
	addr.SetPort(port)
	clientRecieverConn, err := net.ListenUDP("udp4",addr.GetIP())
	if err!=nil{
		fmt.Println("server reciever error")
	}
	go func(){
		for{
			clientAddr,encData:=readfromClient(clientRecieverConn)
			if encData=="INITCONN"{
				clientUDPAddr, err := net.ResolveUDPAddr("udp4", clientAddr)
				if err!=nil{
					fmt.Println(err)
				}
				conn, err := net.DialUDP("udp4", nil, clientUDPAddr)
				if err!=nil{
					fmt.Println(err)
				}
				conn.Write([]byte(addr.GetIP().String()))
				conn.Close()
			}
		}
	}()
}

// func getRandomToken(length int) string {
// 	randomBytes := make([]byte, length)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
// }

func readfromClient(client *net.UDPConn) (string,string) {
	var buf [512]byte
	n, _, err := client.ReadFromUDP(buf[0:])
	if err != nil {
		panic(err)
	}
	str:=strings.Split(string(buf[:n]), ":")
	return str[1],str[2]
}