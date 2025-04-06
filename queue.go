package threadsafe

import (
	"sync"
)

// Queue is a thread-safe queue.
type Queue[T any] struct {
	data []T
	mu   sync.Mutex
}

// NewQueue creates a new thread-safe queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: []T{}}
}

// Enqueue adds an element to the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, value)
}

// Dequeue removes and returns an element from the queue.
func (q *Queue[T]) Dequeue() (value T, ok bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	n := len(q.data)
	if n > 0 {
		value = q.data[0]
		q.data = q.data[1:]
		ok = true
	}
	return
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.data)
}

// Peek returns the element at the front of the queue without removing it.
// Example:
//
//	value, ok := q.Peek()
func (q *Queue[T]) Peek() (value T, ok bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	n := len(q.data)
	if n > 0 {
		value = q.data[0]
		ok = true
	}
	return
}

// IsEmpty checks if the queue is empty.
// Example:
//
//	isEmpty := q.IsEmpty()
func (q *Queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.data) == 0
}

// Clear removes all elements from the queue.
// Example:
//
//	q.Clear()
func (q *Queue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = []T{}
}

// Values returns a slice of all elements in the queue.
// Example:
//
//	values := q.Values()
func (q *Queue[T]) Values() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	values := make([]T, len(q.data))

	copy(values, q.data)

	return values
}
