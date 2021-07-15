package queue

// Queue that stores integers
type Queue struct {
	elements []int
}

// Push adds an elements to the stack
func (q *Queue) Enqueue(i int) {
	q.elements = append(q.elements, i)
}

// Pop removes the first element from the stack and returns it
func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	popped := q.elements[0]
	q.elements = q.elements[1:]

	return popped, true
}

// Read returns the element on the top of the stack
func (q *Queue) Read() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	return q.elements[len(q.elements)-1], true
}

// Empty checks if the queue is empty or not
func (q *Queue) Empty() bool {
	if len(q.elements) == 0 {
		return true
	}

	return false
}
