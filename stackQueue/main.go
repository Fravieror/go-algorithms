package main

import "fmt"

func main() {
	// Stack example
	fmt.Println("Stack Example:")
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack after pushing 1, 2, 3:", stack.elements)

	top, _ := stack.Peek()
	fmt.Println("Peek top element:", top)

	popped, _ := stack.Pop()
	fmt.Println("Popped element:", popped)
	fmt.Println("Stack after pop:", stack.elements)

	isEmpty := stack.IsEmpty()
	fmt.Println("Is stack empty?", isEmpty)

	// Queue example
	fmt.Println("\nQueue Example:")
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("Queue after enqueuing 1, 2, 3:", queue.elements)

	front, _ := queue.Peek()
	fmt.Println("Peek front element:", front)

	dequeued, _ := queue.Dequeue()
	fmt.Println("Dequeued element:", dequeued)
	fmt.Println("Queue after dequeue:", queue.elements)

	isEmpty = queue.IsEmpty()
	fmt.Println("Is queue empty?", isEmpty)
}
