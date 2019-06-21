package piscine

func SortWordArr(array []string) {
	ech:=""
	for i:=0;i<len(array);i++{
		for j:=i+1;j<len(array);j++{
			if []byte(array[j])[0]<[]byte(array[i])[0]{
				ech=array[i]
				array[i]=array[j]
				array[j]=ech
			}
		}
	}
}

