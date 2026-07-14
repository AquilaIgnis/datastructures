package tests

import (
	"math/rand"
	"slices"
	"testing"

	"datastructures/trees/heap"
)

func TestPQPopEmpty(t *testing.T) {
	q := heap.NewPriorityQueue[int]()
	if v, ok := q.Pop(); ok {
		t.Fatalf("Pop on empty queue = (%d, true), want (0, false)", v)
	}
}

func TestPQPeekEmpty(t *testing.T) {
	q := heap.NewPriorityQueue[int]()
	if v, ok := q.Peek(); ok {
		t.Fatalf("Peek on empty queue = (%d, true), want (0, false)", v)
	}
}

func TestPQSizeTracksPushPop(t *testing.T) {
	q := heap.NewPriorityQueue[int]()
	if q.Size() != 0 {
		t.Fatalf("new queue Size = %d, want 0", q.Size())
	}
	for i, v := range []int{9, 4, 7, 1} {
		q.Push(v)
		if q.Size() != i+1 {
			t.Fatalf("after %d pushes Size = %d, want %d", i+1, q.Size(), i+1)
		}
	}
	for want := 3; want >= 0; want-- {
		q.Pop()
		if q.Size() != want {
			t.Fatalf("Size = %d, want %d", q.Size(), want)
		}
	}
}

func TestPQPeekReturnsMinWithoutRemoving(t *testing.T) {
	q := heap.NewPriorityQueue[int]()
	for _, v := range []int{8, 3, 5, 1, 9} {
		q.Push(v)
	}
	v, ok := q.Peek()
	if !ok || v != 1 {
		t.Fatalf("Peek = (%d, %v), want (1, true)", v, ok)
	}
	// Peek must not mutate the queue.
	if q.Size() != 5 {
		t.Fatalf("Size after Peek = %d, want 5", q.Size())
	}
	v2, _ := q.Peek()
	if v2 != v {
		t.Fatalf("second Peek = %d, want same %d", v2, v)
	}
}

func TestPQPopsInPriorityOrder(t *testing.T) {
	in := []int{5, 3, 8, 1, 9, 2, 7, 0, 4, 6}
	q := heap.NewPriorityQueue[int]()
	for _, v := range in {
		q.Push(v)
	}

	got := []int{}
	for q.Size() > 0 {
		// Peek must always agree with the next Pop.
		p, ok := q.Peek()
		if !ok {
			t.Fatal("Peek returned ok=false while Size > 0")
		}
		v, ok := q.Pop()
		if !ok {
			t.Fatal("Pop returned ok=false while Size > 0")
		}
		if p != v {
			t.Fatalf("Peek returned %d but Pop returned %d", p, v)
		}
		got = append(got, v)
	}

	want := slices.Clone(in)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Fatalf("pop order = %v, want ascending %v", got, want)
	}
}

func TestPQStrings(t *testing.T) {
	q := heap.NewPriorityQueue[string]()
	for _, s := range []string{"pear", "apple", "cherry", "banana"} {
		q.Push(s)
	}
	got := []string{}
	for q.Size() > 0 {
		v, _ := q.Pop()
		got = append(got, v)
	}
	want := []string{"apple", "banana", "cherry", "pear"}
	if !slices.Equal(got, want) {
		t.Fatalf("pop order = %v, want %v", got, want)
	}
}

func TestPQFuzzAgainstSort(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(50)
		in := make([]int, n)
		for i := range in {
			in[i] = rng.Intn(100)
		}

		q := heap.NewPriorityQueue[int]()
		for _, v := range in {
			q.Push(v)
		}

		got := []int{}
		for q.Size() > 0 {
			v, _ := q.Pop()
			got = append(got, v)
		}

		want := slices.Clone(in)
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Fatalf("trial %d: got %v, want %v (input %v)", trial, got, want, in)
		}
	}
}
