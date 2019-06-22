package piscine

func Rot14(str string) string {
	bstr:=[]byte(str)
	for i:= range bstr{
		fmt.Println("avant",bstr[i])
		if bstr[i]>=65 && bstr[i]<=90{
			if (bstr[i]+14)>90{
				bstr[i]=bstr[i]-12			
			}else{
				bstr[i]+=14
			}
		}
		if bstr[i]>=97 && bstr[i]<=122{
			if (bstr[i]+14)>122{
				bstr[i]=bstr[i]-12			
			}else{
				bstr[i]+=14
			}
		}
		fmt.Println("apres",bstr[i])
	}
	return string(bstr)
}

