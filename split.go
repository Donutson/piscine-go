package piscine

import "strings"

func Split(str, charset string) []string {
	/*var result []string
	var contain []byte
	str=Concat(str,charset)
	i:=0
	for i<len(str){
		if !ischarset(i , str, charset){
			contain=append(contain,str[i])
		}
		if str[i]==charset[0]{
			if ischarset(i,str,charset){
				result=append(result, string(contain))
				contain=[]byte{}
				i+=len(charset)-1
			}
		}
		i++
	}*/
	return strings.Split(str, charset)
}

func ischarset(i int, str, charset string)bool{
	if len(charset)< len(str[i:]){
		for j,_:= range charset{
			if charset[j]!=str[i]{
				return false
			}
			i++
		}
	}
	
	return true
}


