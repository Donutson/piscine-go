package piscine

func ListReverse(l *List) {
	cl:=&List{}
	if l.Head!=nil{
		cl.Tail=l.Head
		for l.Head.Next!=nil{
			ListPushFront(cl, l.Head.Data)
			l.Head=l.Head.Next
		}
		ListPushFront(cl, l.Head.Data)
		l.Head=cl.Head
		l.Tail=cl.Tail
		l.Tail.Next=nil
	}
}
