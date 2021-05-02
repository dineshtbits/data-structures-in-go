package stack

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

func (s *Stack) Show() []int {
	return s.items
}