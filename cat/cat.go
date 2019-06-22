package main

import (
    "fmt"
    "io/ioutil"
    "os"
)



func read(filename string) (string, error, bool) {
    data, err := ioutil.ReadFile(filename)
    return string(data), err, true
}

func main() {
	filename:=os.Args
	if len(filename)==1{
		fmt.Print()
	}else{
		for i:=1;i<len(filename);i++{
			data,e,fa:=read(filename[1])
			fmt.Println(data)
		}
		if fa{
			fmt.Print(e)		
		}
		
	}
}
