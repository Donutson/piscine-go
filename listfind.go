package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l.Head!=nil{
		cl:=l.Head
		for cl.Next!=nil{
			if comp(cl.Data,ref){
				return &cl.Data
			}
			cl=cl.Next
		}
		if comp(cl.Data,ref){
			return &cl.Data
		}
	}
	return nil
}
