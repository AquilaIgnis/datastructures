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

// MaxHeap :: binary heap implementation,  Insertion - O(log n).
// Array[] is managed internally , treat it as read only
type MaxHeap[T cmp.Ordered] struct {
	Array []T
}

// NewMinHeap() :: Creates a new max heap tree.
func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{}
}

// Insert() -> adds new value to heap.
// O(logn)
func (h *MaxHeap[T]) Insert(val T) {
	h.Array = append(h.Array, val)

	h.siftUp(len(h.Array) - 1)
}

// siftUp() :: moves the larger value node up, used when inserting.
// O(logn)
func (h *MaxHeap[T]) siftUp(i int) {
	for i > 0 {
		parentIndex := (i - 1) / 2

		if h.Array[i] > h.Array[parentIndex] {
			// swap places
			h.Array[i], h.Array[parentIndex] = h.Array[parentIndex], h.Array[i]
			i = parentIndex
		} else {
			break
		}
	}
}

// siftDown() :: receives an index to move its value down the tree.
// O(logn)
func (h *MaxHeap[T]) siftDown(i int) {
	n := len(h.Array)

	for {

		smallest := i

		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.Array[left] > h.Array[smallest] {
			smallest = left
		}
		if right < n && h.Array[right] > h.Array[smallest] {
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

// PopMax() -> returns the root
func (h *MaxHeap[T]) PopMax() (T, bool) {
	var zero T
	if len(h.Array) <= 0 {
		return zero, false
	}

	last := len(h.Array) - 1
	root := h.Array[0]

	// move last value to the root
	h.Array[0] = h.Array[last]

	// GC cleanup
	h.Array[last] = zero

	h.Array = h.Array[:last]

	h.siftDown(0)

	return root, true
}

// Heapify() -> makes arbitrary slice a heap from scratch . returns error if
// the heap is not empty
func (h *MaxHeap[T]) Heapify(array []T) error {
	if len(h.Array) != 0 {
		return errors.New("Can only init Heapify on empty heap")
	}

	h.Array = slices.Clone(array)

	for i := len(h.Array)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return nil
}
