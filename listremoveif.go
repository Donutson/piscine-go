package piscine


func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head!=nil{
		cl:=l.Head
		for cl.Next!=nil{
			if cl.Next.Data==data_ref{
				cl.Next=cl.Next.Next
			}
			cl=cl.Next
		}
		if cl.Next.Data==data_ref{
			cl.Next=cl.Next.Next
		}
	}
}

