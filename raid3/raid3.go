package main

import (
	"fmt"
	"os"
	"io/ioutil"
	)

func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func read(filename string) string {
    data, err := ioutil.ReadFile(filename)
    check(err)
    return string(data)
}

func main(){
	
		data:=read("raid1a.go")
		fmt.Println(data)
	
}

