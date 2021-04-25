package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(element int) {
	s.items = append(s.items, element)
}

func (s *Stack) Pop() int {
	length := len(s.items) - 1
	removed := s.items[length]
	s.items = s.items[:length]
	return removed
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	s.Push(6)
	fmt.Printf("stack is %v\n", s)
	fmt.Printf("removed %v - %v\n", s.Pop(), s.items)
	fmt.Printf("removed %v - %v\n", s.Pop(), s.items)
	fmt.Printf("removed %v - %v\n", s.Pop(), s.items)
	fmt.Printf("removed %v - %v\n", s.Pop(), s.items)
}
