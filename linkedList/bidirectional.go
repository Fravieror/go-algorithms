package main

import "fmt"

type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
}

type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
}

// Insert inserts a new value at the end of the doubly linked list.
func (l *DoublyLinkedList) Insert(value int) {
	newNode := &DNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}
}

// Delete deletes the first occurrence of the value in the doubly linked list.
func (l *DoublyLinkedList) Delete(value int) {
	current := l.Head
	for current != nil && current.Value != value {
		current = current.Next
	}

	if current != nil {
		if current.Prev != nil {
			current.Prev.Next = current.Next
		} else {
			l.Head = current.Next
		}

		if current.Next != nil {
			current.Next.Prev = current.Prev
		} else {
			l.Tail = current.Prev
		}
	}
}

// Search searches for a value in the doubly linked list.
func (l *DoublyLinkedList) Search(value int) *DNode {
	current := l.Head
	for current != nil && current.Value != value {
		current = current.Next
	}
	return current
}

// Reverse reverses the doubly linked list.
func (l *DoublyLinkedList) Reverse() {
	var prev, next *DNode
	current := l.Head
	l.Tail = l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}
	l.Head = prev
}

// Print prints the doubly linked list.
func (l *DoublyLinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}
