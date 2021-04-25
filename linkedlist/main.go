package main

import (
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

type List struct {
	head *Node
}

func (L *List) insertAtBegining(s string) {
	n := &Node{value: s}

	if L.head != nil {
		n.next = L.head
		L.head = n
	} else {
		L.head = n
	}
}

func (L *List) show() {
	currentNode := L.head
	if currentNode == nil {
		fmt.Println("List is empty")
	} else {
		for currentNode != nil {
			fmt.Printf("value: %v, next: %v \n", currentNode.value, currentNode.next)
			currentNode = currentNode.next
		}
	}
}

func main() {
	l := List{}
	l.insertAtBegining("dinesh")
	l.insertAtBegining("Tummlapalli")
	l.show()
}
