package main

import "fmt"

// Insert inserts a new value at the end of the linked list.
func (l *LinkedList) Insert(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// Delete deletes the first occurrence of the value in the linked list.
func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

// Search searches for a value in the linked list.
func (l *LinkedList) Search(value int) *Node {
	current := l.Head
	for current != nil && current.Value != value {
		current = current.Next
	}
	return current
}

// Reverse reverses the linked list.
func (l *LinkedList) Reverse() {
	var prev, next *Node
	current := l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

// Print prints the linked list.
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}
