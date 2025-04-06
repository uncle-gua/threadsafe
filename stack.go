package threadsafe

import (
	"sync"
)

// Stack is a thread-safe stack.
type Stack[T any] struct {
	data []T
	mu   sync.Mutex
}

// NewStack creates a new thread-safe stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, value)
}

// Pop removes and returns an element from the stack.
func (s *Stack[T]) Pop() (value T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := len(s.data)
	if n > 0 {
		value = s.data[n-1]
		s.data = s.data[:n-1]
		ok = true
	}
	return
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}

// Peek returns the element at the top of the stack without removing it.
// Example:
//
//	value, ok := s.Peek()
func (s *Stack[T]) Peek() (value T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := len(s.data)
	if n > 0 {
		value = s.data[n-1]
		ok = true
	}
	return
}

// IsEmpty checks if the stack is empty.
// Example:
//
//	isEmpty := s.IsEmpty()
func (s *Stack[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data) == 0
}

// Clear removes all elements from the stack.
// Example:
//
//	s.Clear()
func (s *Stack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = []T{}
}

// Values returns a slice of all elements in the stack.
// Example:
//
//	values := s.Values()
func (s *Stack[T]) Values() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	n := len(s.data)
	values := make([]T, n)

	for i, value := range s.data {
		values[n-i-1] = value
	}

	return values
}
