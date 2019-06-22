package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        fmt.Print(e)
    }
}

func read(filename string) string {
    data, err := ioutil.ReadFile(filename)
    check(err)
    return string(data)
}

func main() {
	filename:=os.Args
	if len(filename)==1{
		fmt.Print()
	}else{
		for i:=1;i<len(filename);i++{
			data:=read(filename[1])
			fmt.Println(data)
		}
		
	}
}
