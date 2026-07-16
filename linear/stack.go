package linear

import (
	"errors"
)

// Stack -> creates a Stack with no size
type Stack[T any] struct {
	capacity  int
	container []T
}

// NewStackFixed -> creates Stack with a fixed size
func NewStackFixed[T any](value int) *Stack[T] {
	return &Stack[T]{
		capacity:  value,
		container: make([]T, 0, value),
	}
}

// Push -> Appends value to stack returns error
// If using StackFixed errors when over capacity
func (s *Stack[T]) Push(value T) error {
	if s.capacity > 0 && (len(s.container) >= s.capacity) {
		return errors.New("Error max capasity exeded")
	}

	s.container = append(s.container, value)
	return nil
}

// Pop -> returns value,error and removes last index
func (s *Stack[T]) Pop() (T, error) {
	var zero T

	if len(s.container) <= 0 {
		return zero, errors.New("Error Empty Stack")
	}

	last := s.container[len(s.container)-1]

	// GC cleanup
	s.container[len(s.container)-1] = zero

	s.container = s.container[:len(s.container)-1]

	return last, nil
}

// Clear -> clears the stack
func (s *Stack[T]) Clear() {
	s.container = []T{}
}

// Peek -> returns top item
func (s *Stack[T]) Peek() T {
	return s.container[len(s.container)-1]
}

// First -> returns bottom item
func (s *Stack[T]) First() T {
	return s.container[0]
}

// Size -> returns number of items in stack
func (s *Stack[T]) Size() int {
	return len(s.container)
}
