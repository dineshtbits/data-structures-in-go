package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type SinglyList struct {
	head *Node
}

type DoublyList struct {
	head *Node
}

func (l *SinglyList) Show() {
	fmt.Println()
	current := l.head
	for current != nil {
		fmt.Printf("%v --> ", current.data)
		current = current.next
	}
}

func (l *SinglyList) Insert(data int) {
	// Insest at the beginnig of the list
	n := &Node{data: data}
	if l.head != nil {
		n.next = l.head
	}
	l.head = n
}

func (l *SinglyList) InsertAt(position, data int) {
	// Insest at a position in the list
	previous := l.head
	previousPos := 0
	for previous != nil && previousPos != position {
		previousPos++
		previous = previous.next
	}

	if previous == nil {
		fmt.Println("Not able to insert, unknown position")
		return
	}

	n := &Node{data: data}
	n.next = previous.next
	previous.next = n
}

func (l *SinglyList) Append(data int) {
	// Insest at end of the list
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{data: data}
}

func (l *SinglyList) Delete() {
	// Delete head node of the list
	if l.head != nil {
		l.head = l.head.next
	} else {
		fmt.Println("List is empty, nothing to delete")
	}
}

func (l *SinglyList) Pop() {
	// Deletes last node
	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (l *SinglyList) DeleteAt(position int) {
	// Delete node at position of the list
	current := l.head
	currentPos := 0
	for current != nil && currentPos != position-1 {
		currentPos++
		current = current.next
	}

	if current == nil {
		fmt.Println("Position not founc")
		return
	}

	current.next = current.next.next
}

func (l *SinglyList) FindNthNodeFromEnd(n int) int {
	// If length of the list is known we can find the element at len(list)-n+1
	// An alternative is to use two pointers.
	if l.head == nil {
		fmt.Println("List is empty - aborting")
		return 0
	}
	var slowPtr *Node
	current := l.head
	c := 1
	for current != nil && c <= n {
		current = current.next
		c++
	}
	if current == nil {
		fmt.Println("Position not found - aborting")
		return 0
	}

	slowPtr = l.head
	for current != nil {
		current = current.next
		slowPtr = slowPtr.next
	}
	return slowPtr.data
}

func (l *SinglyList) CheckIfLoopsExists() bool {
	slowPtr, fastPtr := l.head, l.head
	for fastPtr != nil && fastPtr.next != nil && slowPtr != nil {
		fastPtr = fastPtr.next.next
		slowPtr = slowPtr.next
		if slowPtr == fastPtr {
			return true
		}
	}
	return false
}

func (l *SinglyList) CreateLoop() {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = l.head
}

func (l *DoublyList) Insert(data, position int) {
	n := &Node{data: data}
	if position == 0 { // insert at the beginnig of the list
		if l.head != nil {
			l.head.prev = n
			n.next = l.head
		}
		l.head = n
	} else { // insert at position
		current := l.head
		currentPos := 0
		for current != nil && currentPos != position { //reach the position
			currentPos++
			current = current.next
		}

		if current == nil {
			fmt.Println("Unable to find position - aborting insert")
			return
		}

		n.next = current.next
		n.prev = current
		if current.next != nil {
			current.next.prev = n
		}
		current.next = n
	}
}

func (l *DoublyList) Delete(position int) {
	if position == 0 { // delete at the beginnig of the list
		if l.head == nil {
			fmt.Println("list is empty - aborting delete")
			return
		}
		l.head.next.prev = nil
		l.head = l.head.next

	} else { // delete at position
		current := l.head
		currentPos := 0
		for current != nil && currentPos != position-1 { //reach the position
			currentPos++
			current = current.next
		}

		if current == nil {
			fmt.Println("Unable to find position - aborting insert")
			return
		}

		if current.next.next != nil {
			temp := current.next.next
			temp.prev = current
			current.next = temp
		} else { // deleting last elemenet
			current.next = nil
		}

	}
}

func (l *DoublyList) Show() {
	fmt.Println()
	current := l.head
	for current != nil {
		fmt.Printf("|*%v|%v|*%v| --> ", nullOrValue(current.prev), nullOrValue(current), nullOrValue(current.next))
		current = current.next
	}
}

func nullOrValue(n *Node) int {
	if n != nil {
		return n.data
	} else {
		return 0
	}
}
