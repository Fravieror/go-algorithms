package main

import "fmt"

func main() {
	// Create a new stack
	stack := &Stack{}

	// Push some contexts onto the stack
	stack.Push("Context 1")
	stack.Push("Context 2")
	stack.Push("Context 3")

	fmt.Println("Stack after pushing Context 1, 2, 3:", stack.elements)

	// Peek the top context
	top, _ := stack.Peek()
	fmt.Println("Peek top context:", top)

	// Pop the top context
	popped, _ := stack.Pop()
	fmt.Println("Popped context:", popped)
	fmt.Println("Stack after pop:", stack.elements)

	// Check if the stack is empty
	isEmpty := stack.IsEmpty()
	fmt.Println("Is stack empty?", isEmpty)

	// Get the size of the stack
	size := stack.Size()
	fmt.Println("Size of stack:", size)

	// Pop remaining contexts
	for !stack.IsEmpty() {
		context, _ := stack.Pop()
		fmt.Println("Popped context:", context)
	}
	fmt.Println("Stack after popping all elements:", stack.elements)
}

// Stack represents a stack data structure
type Stack struct {
	elements []interface{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}
	index := len(s.elements) - 1
	value := s.elements[index]
	s.elements = s.elements[:index]
	return value, true
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}
