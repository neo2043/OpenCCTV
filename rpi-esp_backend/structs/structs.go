package structs

import (
	"fmt"
	"net"
	utils "backend/utils"
)

type FileSaveData struct {
	Name    string `json:"name"`
	UrlPath string `json:"urlpath"`
}

type TempSaveData struct {
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	UrlPath   string `json:"urlpath"`
	Connected bool   `json:"connected"`
}

type TempSaveDataMap map[string]TempSaveData
type FileSaveDataMap map[string]FileSaveData

func (t *TempSaveData) UpdateIP(ip string) TempSaveData {
	return TempSaveData{
		Ip:        ip,
		Name:      t.Name,
		UrlPath:   t.UrlPath,
		Connected: t.Connected,
	}
}

func (t *TempSaveData) UpdateName(name string) TempSaveData {
	return TempSaveData{
		Ip:        t.Ip,
		Name:      name,
		UrlPath:   t.UrlPath,
		Connected: t.Connected,
	}
}

func (t *TempSaveData) UpdateUrlPath(urlpath string) TempSaveData {
	return TempSaveData{
		Ip:        t.Ip,
		Name:      t.Name,
		UrlPath:   urlpath,
		Connected: t.Connected,
	}
}

func (t *TempSaveData) UpdateConnected(connected bool) TempSaveData {
	return TempSaveData{
		Ip:        t.Ip,
		Name:      t.Name,
		UrlPath:   t.UrlPath,
		Connected: connected,
	}
}

func GetTempSaveDataStructMap() TempSaveDataMap {
	return make(TempSaveDataMap)
}

func GetFileSaveDataStructMap() FileSaveDataMap {
	return make(FileSaveDataMap)
}

func (t TempSaveDataMap) UpdateTempSaveDataMap(ip, name, urlpath, connected interface{}, key string) {
	tempStruct := t[key]
	if ip != nil {
		tempStruct = tempStruct.UpdateIP(ip.(string))
	}
	if name != nil {
		tempStruct = tempStruct.UpdateName(name.(string))
	}
	if urlpath != nil {
		tempStruct = tempStruct.UpdateUrlPath(urlpath.(string))
	}
	if connected != nil {
		tempStruct = tempStruct.UpdateConnected(connected.(bool))
	}
	t[key] = tempStruct
}

func keyMap[K string, V interface{}](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func ConvTempSaveDataMaptoFileSaveDataMap(data TempSaveDataMap) FileSaveDataMap {
	tempKey := keyMap(data)
	tempReturnstruct := make(FileSaveDataMap)
	for _, j := range tempKey {
		tempReturnstruct[j] = FileSaveData{Name: data[j].Name, UrlPath: data[j].UrlPath}
	}
	return tempReturnstruct
}

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