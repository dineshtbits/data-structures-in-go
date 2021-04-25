package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(element int) {
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() int {
	dequed := q.items[0]
	q.items = q.items[1:]
	return dequed
}

func main() {
	q := &Queue{}
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
