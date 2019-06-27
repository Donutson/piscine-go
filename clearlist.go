package piscine

import "fmt"

func ListClear(l *List) {
	if l.Head!=nil{
		for l.Head.Next!=nil{
			old:=l.Head
			l.Head=l.Head.Next
			old.Next=nil
		}
	}
}

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

