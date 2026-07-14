package tests

import (
	"math/rand"
	"slices"
	"testing"

	"datastructures/trees/heap"
)

// isMinHeap reports whether arr satisfies the min-heap property: every parent
// is <= its children.
func isMinHeap(arr []int) bool {
	for i := range arr {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(arr) && arr[left] < arr[i] {
			return false
		}
		if right < len(arr) && arr[right] < arr[i] {
			return false
		}
	}
	return true
}

func TestPopMinEmpty(t *testing.T) {
	h := heap.NewMinHeap[int]()
	if v, ok := h.PopMin(); ok {
		t.Fatalf("PopMin on empty heap returned (%d, true), want (0, false)", v)
	}
}

func TestInsertMaintainsHeapProperty(t *testing.T) {
	h := heap.NewMinHeap[int]()
	for _, v := range []int{5, 3, 8, 1, 9, 2, 7, 0, 4, 6} {
		h.Insert(v)
		if !isMinHeap(h.Array) {
			t.Fatalf("heap property violated after inserting %d: %v", v, h.Array)
		}
	}
}

func TestInsertThenPopIsSorted(t *testing.T) {
	in := []int{5, 3, 8, 1, 9, 2, 7, 0, 4, 6}
	h := heap.NewMinHeap[int]()
	for _, v := range in {
		h.Insert(v)
	}

	got := []int{}
	for {
		v, ok := h.PopMin()
		if !ok {
			break
		}
		if !isMinHeap(h.Array) {
			t.Fatalf("heap property violated after popping %d: %v", v, h.Array)
		}
		got = append(got, v)
	}

	want := slices.Clone(in)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Fatalf("popped order = %v, want ascending %v", got, want)
	}
}

func TestPeekRootIsMinimum(t *testing.T) {
	h := heap.NewMinHeap[int]()
	for _, v := range []int{42, 17, 99, 3, 58} {
		h.Insert(v)
		if h.Array[0] != slices.Min(h.Array) {
			t.Fatalf("root = %d, want min %d (heap: %v)", h.Array[0], slices.Min(h.Array), h.Array)
		}
	}
}

func TestDuplicates(t *testing.T) {
	in := []int{4, 4, 1, 1, 4, 1, 2, 2}
	h := heap.NewMinHeap[int]()
	for _, v := range in {
		h.Insert(v)
	}
	got := []int{}
	for {
		v, ok := h.PopMin()
		if !ok {
			break
		}
		got = append(got, v)
	}
	want := slices.Clone(in)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Fatalf("popped order = %v, want %v", got, want)
	}
}

func TestHeapify(t *testing.T) {
	in := []int{9, 4, 7, 1, 0, 3, 8, 2, 6, 5}
	h := heap.NewMinHeap[int]()
	if err := h.Heapify(in); err != nil {
		t.Fatalf("Heapify returned unexpected error: %v", err)
	}
	if !isMinHeap(h.Array) {
		t.Fatalf("Heapify did not produce a valid heap: %v", h.Array)
	}

	// Heapify must clone: mutating the source afterwards must not affect the heap.
	in[0] = -100
	if slices.Contains(h.Array, -100) {
		t.Fatal("Heapify did not clone its input; heap was mutated by source change")
	}

	got := []int{}
	for {
		v, ok := h.PopMin()
		if !ok {
			break
		}
		got = append(got, v)
	}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !slices.Equal(got, want) {
		t.Fatalf("popped order = %v, want %v", got, want)
	}
}

func TestHeapifyOnNonEmptyErrors(t *testing.T) {
	h := heap.NewMinHeap[int]()
	h.Insert(1)
	if err := h.Heapify([]int{2, 3}); err == nil {
		t.Fatal("Heapify on non-empty heap returned nil error, want error")
	}
}

func TestFuzzAgainstSort(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(50)
		in := make([]int, n)
		for i := range in {
			in[i] = rng.Intn(100)
		}

		h := heap.NewMinHeap[int]()
		for _, v := range in {
			h.Insert(v)
		}

		got := []int{}
		for {
			v, ok := h.PopMin()
			if !ok {
				break
			}
			got = append(got, v)
		}

		want := slices.Clone(in)
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Fatalf("trial %d: got %v, want %v (input %v)", trial, got, want, in)
		}
	}
}
