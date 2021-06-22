package advancedqueue

import (
	"errors"
	"fmt"

	"github.com/dineshtbits/data-structures-in-go/stack"
)

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *Queue) Show() {
	fmt.Printf("Size: %v | ", q.size)
	front := q.front
	for front != nil {
		fmt.Printf("%v <--", front.data)
		front = front.next
	}
	fmt.Println()
}

func (q *Queue) Enqueue(data int) {
	n := &Node{data: data}
	if q.rear != nil {
		q.rear.next = n
	}
	q.rear = n
	if q.front == nil {
		q.front = q.rear
	}
	q.size++
}

func (q *Queue) Dequeue() (int, error) {
	if q.size <= 0 {
		return 0, errors.New("Queue is empty")
	}
	temp := q.front
	q.front = q.front.next
	q.size--
	return temp.data, nil
}

func (q *Queue) ReverstFirstNElements(n int) {
	if n > q.size {
		n = q.size
	}

	if n < 1 {
		fmt.Println("Empty queue")
		return
	}

	s := stack.Stack{}
	for i := 1; i <= n; i++ {
		ele, err := q.Dequeue()
		if err == nil {
			s.Push(ele)
		}
	}

	for s.Size() > 0 {
		q.Enqueue(s.Pop().(int))
	}
	for i := 1; i <= (q.size - n); i++ {
		ele, err := q.Dequeue()
		if err == nil {
			q.Enqueue(ele)
		}
	}
}
