package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) bool {
    if e != nil {
        fmt.Print(e)
	return true
    }
	return false
}

func read(filename string) (string, bool) {
    data, err := ioutil.ReadFile(filename)
    a:=check(err)
    return string(data), a
}

func main() {
	filename:=os.Args
	if len(filename)==1{
		fmt.Print()
	}else{
		
		for i:=1;i<len(filename);i++{
			data, a:=read(filename[1])
			fmt.Println(data)
			if a{
				break			
			}
		}
		
	}
}
