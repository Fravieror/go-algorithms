package main

type Stack struct {
	elements []int
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	index := len(s.elements) - 1
	value := s.elements[index]
	s.elements = s.elements[:index]
	return value, true
}

// Peek returns the top element of the stack without removing it.
func (s *Stack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
