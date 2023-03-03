package main

import (
	"algorithms/m/datastructures"
	"container/heap"
	"fmt"
)

func main() {
	// Priority queue implementation

}

func priorityQueue() {
	pq := make(datastructures.PriorityQueue, 0)
	heap.Init(&pq)

	item1 := &datastructures.Item{Value: "foo", Priority: 3}
	heap.Push(&pq, item1)

	item2 := &datastructures.Item{Value: "bar", Priority: 1}
	heap.Push(&pq, item2)

	item3 := &datastructures.Item{Value: "baz", Priority: 2}
	heap.Push(&pq, item3)

	// Get the items from the priority queue in order of priority
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*datastructures.Item)
		fmt.Printf("%s (priority %d)\n", item.Value, item.Priority)
	}
}
