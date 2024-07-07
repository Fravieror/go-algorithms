package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	fmt.Println("Unidirectional Linked List:")
	uList := &LinkedList{}
	uList.Insert(1)
	uList.Insert(2)
	uList.Insert(3)
	uList.Print() // Output: 1 2 3
	uList.Delete(2)
	uList.Print()                       // Output: 1 3
	fmt.Println(uList.Search(3) != nil) // Output: true
	uList.Reverse()
	uList.Print() // Output: 3 1

	fmt.Println("\nDoubly Linked List:")
	dList := &DoublyLinkedList{}
	dList.Insert(1)
	dList.Insert(2)
	dList.Insert(3)
	dList.Print() // Output: 1 2 3
	dList.Delete(2)
	dList.Print()                       // Output: 1 3
	fmt.Println(dList.Search(3) != nil) // Output: true
	dList.Reverse()
	dList.Print() // Output: 3 1

	fmt.Println("\nCircular Linked List:")
	cList := &CircularLinkedList{}
	cList.Insert(1)
	cList.Insert(2)
	cList.Insert(3)
	cList.Print() // Output: 1 2 3
	cList.Delete(2)
	cList.Print()                       // Output: 1 3
	fmt.Println(cList.Search(3) != nil) // Output: true
}
