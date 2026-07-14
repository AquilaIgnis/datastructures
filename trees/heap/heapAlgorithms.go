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
