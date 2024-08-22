package main

import "fmt"


type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	for node:= l.Head; node != nil; node = node.Next{
		count++
	}
	return count
}


func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data}

	if l.Head == nil{
		l.Head = node
		l.Tail = node
	}  else {
		node.Next = l.Head
		l.Head = node
	}
}


func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")
	ListPushFront(link, "")

	fmt.Println(ListSize(link))

}