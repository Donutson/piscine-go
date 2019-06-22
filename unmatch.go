package piscine

func Unmatch(arr []int) int {
	result:=0
	for i:=range arr{
		if !ispair(arr[i],i, arr){
			result=arr[i]
			return result		
		}
	}
	return result
}

func ispair(elt,index int, arr[]int) bool{
	nb:=0
	for i:=range arr{
		if i!=index{
			if arr[i]==elt{
				nb++			
			}
		}
	}
	if nb==1{
		return true
	}
	return false	
}
