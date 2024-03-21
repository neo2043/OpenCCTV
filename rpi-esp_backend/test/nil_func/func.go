package main

import (
	"fmt"
)

func main(){
	update(nil,"sdcs",nil,false)	
}

func update(ip,name,urlpath,connected interface{}){
	fmt.Println(ip==nil)
	fmt.Println(name)
	fmt.Println(urlpath)
	fmt.Println(connected)
}