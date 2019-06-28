package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2.Head!=nil{
		cl2:=l2.Head
		for cl2!=nil{
			ListPushBack(l1, cl2.Data)
			cl2=cl2.Next
		}
	}
}

