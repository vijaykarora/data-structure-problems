package base

type Stack struct {
	Items []rune
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Push(c rune) {
	s.Items = append(s.Items, c)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(s.Items) - 1
	value := s.Items[index]
	s.Items = s.Items[:index]
	return value, true
}

func (s *Stack) Top() rune {
	if s.IsEmpty() {
		return -1
	}

	index := len(s.Items) - 1
	value := s.Items[index]
	return value
}
