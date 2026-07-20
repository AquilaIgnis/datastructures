package heap

import (
	"cmp"
)

// HeapSrort() -> takes in an array and returns it sorted. O(n log n)
func HeapSrort[T cmp.Ordered](arg []T) []T {
	if len(arg) <= 0 {
		return nil
	}

	heap := NewMinHeap[T]()
	heap.Heapify(arg)

	sorted := make([]T, 0, len(arg))

	for range heap.Array {

		val, ok := heap.PopMin()

		if !ok {
			break
		}

		sorted = append(sorted, val)
	}

	return sorted
}

// PriorityQueue
type PriorityQueue[T cmp.Ordered] struct {
	heap MinHeap[T]
}

// NewPriorityQueue() -> creates a Priority Queue using min Heap
func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

// Push() -> adds new value to queue
func (q *PriorityQueue[T]) Push(v T) {
	q.heap.Insert(v)
}

// Pop() -> removes and returns the min value
func (q *PriorityQueue[T]) Pop() (T, bool) {
	return q.heap.PopMin()
}

// Peek() ->  returns the upcoming item without removing it.
// ok is false when the queue is empty.
func (q *PriorityQueue[T]) Peek() (T, bool) {
	var zero T
	if len(q.heap.Array) == 0 {
		return zero, false
	}
	return q.heap.Array[0], true
}

// Size() -> return the Size of queue
func (q PriorityQueue[T]) Size() int {
	return len(q.heap.Array)
}

// Display() -> return the whole queue
func (q PriorityQueue[T]) Display() []T {
	return q.heap.Array
}
