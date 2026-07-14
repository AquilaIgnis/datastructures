package heap

import (
	"cmp"
	"errors"
	"slices"
)

//  node is stored at index i:
//  left child is at index 2*i + 1.
//  right child is at index 2*i + 2.
//  parent index [(i-1)/2].

// MinHeap :: binary heap implementation,  Insertion - O(log n).
// Array[] is managed internally , treat it as read only
type MinHeap[T cmp.Ordered] struct {
	Array []T
}

// NewMinHeap() :: Creates a new min heap tree.
func NewMinHeap[T cmp.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{}
}

// Insert() -> adds new value to heap.
// O(logn)
func (h *MinHeap[T]) Insert(val T) {
	h.Array = append(h.Array, val)

	h.siftUp(len(h.Array) - 1)
}

// siftUp() :: moves the smaller value node up, used when inserting.
// O(logn)
func (h *MinHeap[T]) siftUp(i int) {
	for i > 0 {
		parentIndex := (i - 1) / 2

		if h.Array[i] < h.Array[parentIndex] {
			// swap places
			h.Array[i], h.Array[parentIndex] = h.Array[parentIndex], h.Array[i]
		}

		i = parentIndex
	}
}

// siftDown() :: receives an index to move its value down the tree.
// O(logn)
func (h *MinHeap[T]) siftDown(i int) {
	n := len(h.Array)

	for {

		smallest := i

		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.Array[left] < h.Array[smallest] {
			smallest = left
		}
		if right < n && h.Array[right] < h.Array[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		// sink down
		h.Array[i], h.Array[smallest] = h.Array[smallest], h.Array[i]
		i = smallest

	}
}

// PopMin() -> returns the root
func (h *MinHeap[T]) PopMin() (T, bool) {
	var zero T
	if len(h.Array) <= 0 {
		return zero, false
	}

	last := len(h.Array) - 1
	root := h.Array[0]

	// swap places with last index
	h.Array[0], h.Array[last] = h.Array[last], h.Array[0]

	h.Array = h.Array[0:last]

	h.siftDown(0)

	// GC cleanup
	h.Array[last] = zero

	return root, true
}

// Heapify() -> makes arbitrary slice a heap from scratch . returns error if
// the heap is not empty
func (h *MinHeap[T]) Heapify(array []T) error {
	if len(h.Array) != 0 {
		return errors.New("Can only init Heapify on empty heap")
	}

	h.Array = slices.Clone(array)

	for i := len(h.Array)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return nil
}
