package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// Min-Heap example
	fmt.Println("Min-Heap Example:")
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("Minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()

	// Priority Queue example
	fmt.Println("Priority Queue Example:")
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%s (%d) ", item.value, item.priority)
	}
	fmt.Println()
}
