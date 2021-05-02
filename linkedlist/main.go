package linkedlist

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

func (L *List) InsertAtBegining(s string) {
	n := &Node{value: s}

	if L.head != nil {
		n.next = L.head
	}
	L.head = n
}

func (L *List) DeleteByValue(s string) {
	if L.head == nil {
		fmt.Println("List is empty - nothing to delete")
		return
	}

	if L.head.value == s {
		L.head = L.head.next
		return
	}

	previousNode := L.head

	for previousNode.next != nil && previousNode.next.value != s {
		previousNode = previousNode.next
	}

	if previousNode.next == nil {
		fmt.Println("Not found")
		return
	}
	previousNode.next = previousNode.next.next
}

func (L *List) Show() {
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