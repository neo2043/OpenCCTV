package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

type fileSaveData struct {
	Name string `json:"name"`
	// Id   string `json:"id"`
}

type tempSaveData struct {
	// Id        string `json:"id"`
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	Connected bool   `json:"connected"`
}

// func parseFromFile(){
// 	var t1 map[string]fileSaveData
// 	var t2 map[string]tempSaveData

// 	path, ok := os.LookupEnv("HOMEPATH")
// 	if ok{
// 		dat, err := os.ReadFile(path+"/OpenCCTV/test")
// 		if err!=nil{
// 			fmt.Println(err)
// 		}
// 		if err:=json.Unmarshal(dat,&t1);err!=nil{
// 			fmt.Println("error")
// 		}

// 	}
// }

// func (t* tempSaveData)setIP(ip string){
// 	t.Ip=ip
// }

// func keyMap[K string, V fileSaveData](m map[K]V) []K {
// 	r := make([]K, 0, len(m))
// 	for k := range m {
// 		r = append(r, k)
// 	}
// 	return r
// }

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	fmt.Println(exPath)

	syscall.Chdir(exPath)

	m2:=make(map[string]tempSaveData)
	m2["sdfgdfbdb"]=tempSaveData{
		Name: "eps1",
		Ip: "192.168.1.1",
		Connected: false,
	}
	m2["sdfgdfbdf"]=tempSaveData{
		Name: "eps1",
		Ip: "192.168.1.1",
		Connected: false,
	}

	m := make(map[string]fileSaveData)
	data,_:=json.Marshal(m2)
	_ = json.Unmarshal(data,&m)
	
	fmt.Println(m)
	
	// // fmt.Println(string(d))
	// value1, _ := os.LookupEnv("HOMEDRIVE")
	// value2, _ := os.LookupEnv("HOMEPATH")
	// err = os.WriteFile(value1+value2+"\\test", data, 0644)
	// fmt.Println(err)

	_, err = os.ReadFile("./test")
	if err != nil {
		fmt.Println(err)
	}
	// if err := json.Unmarshal(dat, &m); err != nil {
	// 	fmt.Println("error")
	// }

	// // fmt.Println(keyMap(m))
	// for _,k := range keyMap(m) {
	// 	m2[k]=tempSaveData{Id: m[k].Id,Name: m[k].Name}
	// }
	// fmt.Println(m2["1"].Connected)
	// fmt.Println(m2["1"].Id)
	// fmt.Println(m2["1"].Ip)
	// fmt.Println(m2["1"].Name)

	// m2["1"].Ip=string("aeffewefwef")
}