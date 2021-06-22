package stack

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.items = append(s.items, element)
}

func (s *Stack) Top() interface{} {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() interface{} {
	length := len(s.items) - 1
	removed := s.items[length]
	s.items = s.items[:length]
	return removed
}

func (s *Stack) Show() []interface{} {
	return s.items
}

func (s *Stack) Size() int {
	return len(s.items)
}
