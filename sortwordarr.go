package piscine

func SortWordArr(array []string) {
	ech:=""
	for i:=0;i<len(array);i++{
		for j:=i+1;j<len(array);j++{
			a:=[]byte(array[j])
			b:=[]byte(array[i])
			if a[0]<b[0]{
				ech=array[i]
				array[i]=array[j]
				array[j]=ech
			}
		}
	}
}

