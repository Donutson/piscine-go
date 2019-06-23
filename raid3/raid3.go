package main

import (
	"fmt"
	"io/ioutil"
	)

func main(){
	param,_:=ioutil.ReadFile("raid1a.go")
	fmt.Println(string(param))
}

