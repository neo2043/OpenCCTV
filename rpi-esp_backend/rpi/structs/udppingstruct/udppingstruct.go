package udppingstruct

import (
	"fmt"
	"net"
	utils "backend/utils"
)

type UdpAddr struct {
	addr *net.UDPAddr
}

func (t *UdpAddr) SetPort(port int) {
	t.addr.Port = port
}

func (t *UdpAddr) SetIP(ip interface{}) error {
	switch argType := ip.(type) {
	case string:
		if t.addr != nil {
			t.addr.IP = net.ParseIP(ip.(string))
		} else {
			t.addr = &net.UDPAddr{}
			t.addr.IP = net.ParseIP(ip.(string))
		}
	case []byte:
		ip := ip.([]byte)
		if t.addr != nil {
			t.addr.IP = net.IPv4(ip[0], ip[1], ip[2], ip[3])
		} else {
			t.addr = &net.UDPAddr{}
			t.addr.IP = net.IPv4(ip[0], ip[1], ip[2], ip[3])
		}
	default:
		return fmt.Errorf("data type not recognized: %T",argType)
	}
	return nil
}

func (t *UdpAddr) GetIP() []byte {
	if t.addr == nil {
		fmt.Println("t.addr is nil")
	}
	return t.addr.IP
}

func (t *UdpAddr) GetPort() string {
	return fmt.Sprintf("%d",t.addr.Port)
}

func (t *UdpAddr) GetIP_Port() string {
	return utils.ByteArraytoString(t.GetIP())+t.GetPort()
}

func (t *UdpAddr) GetUDPAddrStruct() *net.UDPAddr {
	return t.addr
}