package piscine

func ListAt(l *NodeL, pos int) *NodeL{
	if pos<0{
		return nil
	}
	count:=0
	
	if l!=nil{
		cl:=l
		for cl.Next!=nil && count<pos{
			cl=cl.Next
			count++
		}
		if pos-count>0{
			return nil
		}
		return cl
	}
	return nil
}

