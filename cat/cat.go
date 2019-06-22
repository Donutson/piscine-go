package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error)(error,bool) {
   return e, true
}

func read(filename string) (string, error, bool) {
    data, err := ioutil.ReadFile(filename)
    e,fa:=check(err)
    return string(data), e, fa
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
