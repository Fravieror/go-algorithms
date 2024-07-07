package main

import "fmt"

type CNode struct {
	Value int
	Next  *CNode
}

type CircularLinkedList struct {
	Head *CNode
}

// Insert inserts a new value into the circular linked list.
func (l *CircularLinkedList) Insert(value int) {
	newNode := &CNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		newNode.Next = l.Head
	} else {
		current := l.Head
		for current.Next != l.Head {
			current = current.Next
		}
		current.Next = newNode
		newNode.Next = l.Head
	}
}

// Delete deletes the first occurrence of the value in the circular linked list.
func (l *CircularLinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}

	current := l.Head
	var prev *CNode

	for {
		if current.Value == value {
			if prev == nil {
				last := l.Head
				for last.Next != l.Head {
					last = last.Next
				}
				if last == l.Head {
					l.Head = nil
				} else {
					last.Next = l.Head.Next
					l.Head = l.Head.Next
				}
			} else {
				prev.Next = current.Next
			}
			return
		}

		prev = current
		current = current.Next
		if current == l.Head {
			break
		}
	}
}

// Search searches for a value in the circular linked list.
func (l *CircularLinkedList) Search(value int) *CNode {
	if l.Head == nil {
		return nil
	}

	current := l.Head
	for {
		if current.Value == value {
			return current
		}
		current = current.Next
		if current == l.Head {
			break
		}
	}
	return nil
}

// Print prints the circular linked list.
func (l *CircularLinkedList) Print() {
	if l.Head == nil {
		return
	}

	current := l.Head
	for {
		fmt.Print(current.Value, " ")
		current = current.Next
		if current == l.Head {
			break
		}
	}
	fmt.Println()
}
