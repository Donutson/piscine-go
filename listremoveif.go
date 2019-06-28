package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
		for l.Head!=nil{
			if l.Head.Data==data_ref{
				l.Head=l.Head.Next
			}else{
				break
			}		
		}
	if l.Head!=nil{
		old:=l.Head
		cur:=l.Head
		for cur.Next!=nil{
			if cur.Data==data_ref{
				old.Next=cur.Next
			}else{
				old=cur
			}
			cur=cur.Next
		}
		if cur.Data==data_ref{
			old.Next=cur.Next
		}else{
			old=cur
		}
	}
}
