package main

import (
	"fmt"
	"os"
	"strings"
	)
func main(){
	param:=os.Args[1]
	if strings.Index(param,"raid1a")!=-1{
		fmt.Println("[raid1a] [12] [34] || [raid1b] [12] [34] || [raid1c] [12] [34] || [raid1d] [12] [34] || [raid1e] [12] [34]")
	}
}

