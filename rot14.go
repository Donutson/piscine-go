package piscine

func Rot14(str string) string {
	bstr:=[]string{str}
	for i!= range bstr{
		bstr[i]+=14
	}
	return string(bstr)
}

