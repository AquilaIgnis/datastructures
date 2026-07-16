package linear

import (
	"errors"
)

// Queue []Type{} -> creates a Queue with no fixed size
// Implemented as slice
type Queue[T any] struct {
	capacity  uint8
	container []T
}

// NewQueueFixed -> NewQueueFixed[Type](size) creates a Queue with no fixed size
func NewQueueFixed[T any](value uint8) *Queue[T] {
	return &Queue[T]{
		capacity:  value,
		container: make([]T, 0, value),
	}
}

// Push() -> enqeues an item, returns err if queue if full
func (q *Queue[T]) Push(item T) error {
	if q.capacity > 0 && q.capacity <= uint8(len(q.container)) {
		return errors.New("Queue full")
	}

	q.container = append(q.container, item)
	return nil
}

// Peek() -> returns the upcoming item
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.container) <= 0 {
		var zero T
		return zero, false
	}

	return q.container[0], true
}

// Pop() -> returns the upcoming item and removes it from Queue
func (q *Queue[T]) Pop() (T, error) {
	var zero T

	if len(q.container) <= 0 {
		return zero, errors.New("Queue is empty")
	}

	upcoming := q.container[0]
	// if zero is a pointer the GC clears it
	q.container[0] = zero
	q.container = q.container[1:]

	return upcoming, nil
}

// Cull() -> removes the last item , returns it and err
func (q *Queue[T]) Cull() (T, error) {
	if len(q.container) <= 0 {
		var zero T
		return zero, errors.New("Queue is empty")
	}

	last := q.container[len(q.container)-1]
	q.container = q.container[0 : len(q.container)-1]

	return last, nil
}

// Size() -> returns size of queue
func (q *Queue[T]) Size() int {
	return len(q.container)
}
