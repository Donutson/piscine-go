package main

import ("fmt"
	"os"
	"strconv"
	)

func main(){
	if len(os.Args)==4{
		param:=os.Args[1:]
		if !iscorrectfordoop(param){
			fmt.Println(0)	
		}else{
			switch param[1]{
				case "+": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						fmt.Println(a+b)
					}
				case "-": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						fmt.Println(a-b)
					}
				case "*": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						fmt.Println(a*b)
					}
				case "/": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						if b==0{
							fmt.Println("No division by 0")
						}else{
							fmt.Println(a/b)
						}
						
					}
				case "%": {
						a,_:=strconv.Atoi(param[0])
						b,_:=strconv.Atoi(param[2])
						if b==0{
							fmt.Println("No modulo by 0")
						}else{
							fmt.Println(a%b)
						}
						
					}
			}
		}
	}
}

func iscorrectfordoop(param []string) bool{
	if (!IsNumeric(param[0]) || !IsNumeric(param[2])) || (param[1]!=string(byte(43)) && param[1]!=string(byte(42)) && param[1]!=string(byte(47)) && param[1]!=string(byte(37)) && param[1]!=string(byte(45))){
		return false
	}
	return true
}

