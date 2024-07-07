package main

type Queue struct {
	elements []int
}

// Enqueue adds an element to the back of the queue.
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element of the queue.
func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value, true
}

// Peek returns the front element of the queue without removing it.
func (q *Queue) Peek() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	return q.elements[0], true
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
