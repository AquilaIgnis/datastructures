package tests

import (
	"math/rand"
	"sort"
	"testing"

	"datastructures/trees/heap"
)

// isMaxHeap verifies the max-heap property over the backing array: every
// parent is >= each of its children.
func isMaxHeap(a []int) bool {
	for i := range a {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(a) && a[i] < a[left] {
			return false
		}
		if right < len(a) && a[i] < a[right] {
			return false
		}
	}
	return true
}

func TestMaxHeap_PopEmpty(t *testing.T) {
	h := heap.NewMaxHeap[int]()

	val, ok := h.PopMax()
	if ok {
		t.Fatalf("PopMax on empty heap: got ok=true, val=%d; want ok=false", val)
	}
	if val != 0 {
		t.Fatalf("PopMax on empty heap: got val=%d; want zero value 0", val)
	}
}

func TestMaxHeap_InsertMaintainsHeapProperty(t *testing.T) {
	h := heap.NewMaxHeap[int]()

	for _, v := range []int{5, 3, 8, 1, 9, 2, 7, 6, 4, 0} {
		h.Insert(v)
		if !isMaxHeap(h.Array) {
			t.Fatalf("heap property violated after inserting %d: %v", v, h.Array)
		}
	}
}

func TestMaxHeap_PopMaxReturnsDescending(t *testing.T) {
	h := heap.NewMaxHeap[int]()

	input := []int{5, 3, 8, 1, 9, 2, 7, 6, 4, 0}
	for _, v := range input {
		h.Insert(v)
	}

	var got []int
	for {
		val, ok := h.PopMax()
		if !ok {
			break
		}
		got = append(got, val)

		// heap property must hold after every pop
		if !isMaxHeap(h.Array) {
			t.Fatalf("heap property violated after pop: %v", h.Array)
		}
	}

	want := make([]int, len(input))
	copy(want, input)
	sort.Sort(sort.Reverse(sort.IntSlice(want)))

	if len(got) != len(want) {
		t.Fatalf("popped %d values; want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("pop order = %v; want descending %v", got, want)
		}
	}
}

func TestMaxHeap_Duplicates(t *testing.T) {
	h := heap.NewMaxHeap[int]()

	for _, v := range []int{4, 4, 4, 2, 2, 7, 7, 7, 7} {
		h.Insert(v)
	}

	prev := int(^uint(0) >> 1) // max int, so first pop always <= prev
	for {
		val, ok := h.PopMax()
		if !ok {
			break
		}
		if val > prev {
			t.Fatalf("pop order not non-increasing: got %d after %d", val, prev)
		}
		prev = val
	}
}

func TestMaxHeap_Heapify(t *testing.T) {
	h := heap.NewMaxHeap[int]()

	input := []int{5, 3, 8, 1, 9, 2, 7, 6, 4, 0}
	if err := h.Heapify(input); err != nil {
		t.Fatalf("Heapify returned error: %v", err)
	}

	if !isMaxHeap(h.Array) {
		t.Fatalf("Heapify did not produce a valid max heap: %v", h.Array)
	}

	// Heapify must clone: mutating the source must not affect the heap.
	input[0] = 999
	if h.Array[0] == 999 {
		t.Fatalf("Heapify did not clone the input slice")
	}

	// The root must be the maximum.
	if h.Array[0] != 9 {
		t.Fatalf("root after Heapify = %d; want max 9", h.Array[0])
	}
}

func TestMaxHeap_HeapifyOnNonEmptyErrors(t *testing.T) {
	h := heap.NewMaxHeap[int]()
	h.Insert(1)

	if err := h.Heapify([]int{2, 3}); err == nil {
		t.Fatal("Heapify on non-empty heap: got nil error; want error")
	}
}

func TestMaxHeap_HeapifyThenPopSorts(t *testing.T) {
	input := []int{42, -7, 0, 15, 15, 3, 99, -100, 8}

	h := heap.NewMaxHeap[int]()
	if err := h.Heapify(input); err != nil {
		t.Fatalf("Heapify returned error: %v", err)
	}

	var got []int
	for {
		val, ok := h.PopMax()
		if !ok {
			break
		}
		got = append(got, val)
	}

	want := make([]int, len(input))
	copy(want, input)
	sort.Sort(sort.Reverse(sort.IntSlice(want)))

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Heapify+PopMax = %v; want %v", got, want)
		}
	}
}

func TestMaxHeap_Randomized(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for trial := 0; trial < 50; trial++ {
		n := r.Intn(200)
		input := make([]int, n)
		for i := range input {
			input[i] = r.Intn(1000) - 500
		}

		h := heap.NewMaxHeap[int]()
		for _, v := range input {
			h.Insert(v)
		}
		if !isMaxHeap(h.Array) {
			t.Fatalf("trial %d: invalid heap after inserts: %v", trial, h.Array)
		}

		var got []int
		for {
			val, ok := h.PopMax()
			if !ok {
				break
			}
			got = append(got, val)
		}

		want := make([]int, len(input))
		copy(want, input)
		sort.Sort(sort.Reverse(sort.IntSlice(want)))

		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("trial %d: pop order = %v; want %v", trial, got, want)
			}
		}
	}
}
