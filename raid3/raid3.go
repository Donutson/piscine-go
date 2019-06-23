package main

import (
	"fmt"
	"io/ioutil"
	)

func main(){
	param,e:=ioutil.ReadFile("raid1a")
	if e!=nil{
		fmt.Println(e)
	}
	fmt.Println(string(param))
}

