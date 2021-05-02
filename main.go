package main

import (
	"fmt"

	"github.com/dineshtbits/data-structures-in-go/heap"
	"github.com/dineshtbits/data-structures-in-go/linkedlist"
	"github.com/dineshtbits/data-structures-in-go/queue"
	"github.com/dineshtbits/data-structures-in-go/stack"
	"github.com/dineshtbits/data-structures-in-go/tree"
)

func main() {
	// linkedlistFunctions()
	// queueFunctions()
	// stackFunctions()
	// heapFunctions()
	treeFunctions()
}

func treeFunctions() {
	b := &tree.BinaryTree{}
	b.Insert(5)
	b.Insert(20)
	b.Insert(12)
	b.Insert(14)
	b.Insert(15)
	fmt.Printf("b.Max() %v\n", b.Max())
}

func stackFunctions() {
	s := &stack.Stack{}
	s.Push(10)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	s.Push(6)
	fmt.Printf("stack is %v\n", s)
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
}

func heapFunctions() {
	h := &heap.Heap{}
	fmt.Printf("initial value: %v \n", h)
	arr := []int{5, 7, 2, 9, 3, 6, 1, 4, 10}
	for _, v := range arr {
		h.Insert(v)
	}
	fmt.Printf("heapified value: %v \n", h)
	h.Extract()
	fmt.Printf("Extracted value: %v \n", h)
	h.Extract()
	fmt.Printf("Extracted value: %v \n", h)
	h.Extract()
	fmt.Printf("Extracted value: %v \n", h)
}

func linkedlistFunctions() {
	l := linkedlist.List{}
	l.InsertAtBegining("dinesh")
	l.InsertAtBegining("Tummlapalli")
	l.InsertAtBegining("John")
	l.InsertAtBegining("Doe")
	l.Show()
	l.DeleteByValue("David")
	l.Show()
}

func queueFunctions() {
	q := &queue.Queue{}
	q.Enqueue(7)
	q.Enqueue(2)
	q.Enqueue(6)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(87)
	fmt.Printf("Queue : %v\n", q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
}
