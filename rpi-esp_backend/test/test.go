package main

import (
	"encoding/hex"
	"fmt"
)

type dataType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

func main() {
	cborData, _ := hex.DecodeString("a42050231f4c4d4d3051fdc2ec0a3851d5b3830104024c53796d6d6574726963313238030a")
	fmt.Println(cborData)
}
