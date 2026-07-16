package tests

import (
	"testing"

	"datastructures/linear"
)

// --- map type ---

func TestQueueWithMaps(t *testing.T) {
	q := linear.NewQueueFixed[map[string]int](3)

	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"b": 2, "c": 3}
	var nilMap map[string]int

	q.Push(m1)
	q.Push(m2)
	q.Push(nilMap)

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	got, ok := q.Peek()
	if !ok || got["a"] != 1 {
		t.Errorf("Peek returned wrong map: %v", got)
	}

	pulled, err := q.Pop()
	if err != nil || pulled["a"] != 1 {
		t.Errorf("Pull returned wrong map: %v err: %v", pulled, err)
	}

	if q.Size() != 2 {
		t.Errorf("expected size 2 after Pull, got %d", q.Size())
	}
}

func TestQueueMapMutationAfterEnqueue(t *testing.T) {
	q := linear.NewQueueFixed[map[string]int](2)

	m := map[string]int{"x": 10}
	q.Push(m)

	// mutate original map after enqueue — queue holds a reference
	m["x"] = 999

	peeked, _ := q.Peek()
	if peeked["x"] != 999 {
		t.Logf("map mutation not reflected (value copied): got %d", peeked["x"])
	} else {
		t.Logf("map mutation reflected (reference held): got %d", peeked["x"])
	}
}

func TestQueueMapOverCapacity(t *testing.T) {
	q := linear.NewQueueFixed[map[string]int](1)
	q.Push(map[string]int{"a": 1})

	err := q.Push(map[string]int{"b": 2})
	if err == nil {
		t.Error("expected error when adding to full map queue, got nil")
	}
}

// --- struct (object) type ---

type Person struct {
	Name string
	Age  int
}

func TestQueueWithStructs(t *testing.T) {
	q := linear.NewQueueFixed[Person](3)

	q.Push(Person{"Alice", 30})
	q.Push(Person{"Bob", 25})
	q.Push(Person{}) // zero value struct

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	first, err := q.Pop()
	if err != nil || first.Name != "Alice" {
		t.Errorf("expected Alice, got %v", first)
	}

	last, err := q.Cull()
	if err != nil || last.Name != "" {
		t.Errorf("expected zero-value struct at tail, got %v", last)
	}
}

func TestQueueStructOverCapacity(t *testing.T) {
	q := linear.NewQueueFixed[Person](2)
	q.Push(Person{"A", 1})
	q.Push(Person{"B", 2})

	err := q.Push(Person{"C", 3})
	if err == nil {
		t.Error("expected error adding struct beyond capacity, got nil")
	}
}

// --- pointer type ---

func TestQueueWithPointers(t *testing.T) {
	q := linear.NewQueueFixed[*Person](3)

	p1 := &Person{"Alice", 30}
	p2 := &Person{"Bob", 25}

	q.Push(p1)
	q.Push(p2)
	q.Push(nil) // nil pointer

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	got, ok := q.Peek()
	if !ok || got != p1 {
		t.Errorf("Peek: expected p1 pointer, got %v", got)
	}
}

func TestQueueNilPointerPull(t *testing.T) {
	q := linear.NewQueueFixed[*Person](2)
	q.Push(nil)

	got, err := q.Pop()
	if err != nil {
		t.Fatalf("unexpected error pulling nil pointer: %v", err)
	}
	if got != nil {
		t.Errorf("expected nil pointer back, got %v", got)
	}
}

func TestQueuePointerMutationAfterEnqueue(t *testing.T) {
	q := linear.NewQueueFixed[*Person](1)

	p := &Person{"Alice", 30}
	q.Push(p)

	p.Name = "Mutated"

	peeked, _ := q.Peek()
	if peeked.Name != "Mutated" {
		t.Errorf("expected mutation to be reflected via pointer, got %s", peeked.Name)
	}
}

// --- slice type ---

func TestQueueWithSlices(t *testing.T) {
	q := linear.NewQueueFixed[[]int](3)

	q.Push([]int{1, 2, 3})
	q.Push([]int{}) // empty slice
	q.Push(nil)     // nil slice

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	first, err := q.Pop()
	if err != nil || len(first) != 3 || first[0] != 1 {
		t.Errorf("unexpected first slice: %v", first)
	}
}

func TestQueueSliceMutationAfterEnqueue(t *testing.T) {
	q := linear.NewQueueFixed[[]int](1)

	s := []int{1, 2, 3}
	q.Push(s)

	s[0] = 999

	peeked, _ := q.Peek()
	if peeked[0] != 999 {
		t.Logf("slice mutation not reflected (copied): got %d", peeked[0])
	} else {
		t.Logf("slice mutation reflected (shared backing array): got %d", peeked[0])
	}
}

func TestQueueSliceOverCapacity(t *testing.T) {
	q := linear.NewQueueFixed[[]int](1)
	q.Push([]int{1})

	err := q.Push([]int{2})
	if err == nil {
		t.Error("expected error adding slice beyond capacity, got nil")
	}
}

// --- empty queue edge cases across all types ---

func TestQueuePullEmptyMap(t *testing.T) {
	q := linear.NewQueueFixed[map[string]int](2)
	_, err := q.Pop()
	if err == nil {
		t.Error("expected error pulling from empty map queue")
	}
}

func TestQueuePullEmptyPointer(t *testing.T) {
	q := linear.NewQueueFixed[*Person](2)
	_, err := q.Pop()
	if err == nil {
		t.Error("expected error pulling from empty pointer queue")
	}
}

func TestQueuePullEmptySlice(t *testing.T) {
	q := linear.NewQueueFixed[[]int](2)
	_, err := q.Pop()
	if err == nil {
		t.Error("expected error pulling from empty slice queue")
	}
}

func TestQueueCullEmptyStruct(t *testing.T) {
	q := linear.NewQueueFixed[Person](2)
	_, err := q.Cull()
	if err == nil {
		t.Error("expected error culling from empty struct queue")
	}
}

func TestQueuePeekEmptyPointer(t *testing.T) {
	q := linear.NewQueueFixed[*Person](2)
	val, ok := q.Peek()
	if ok || val != nil {
		t.Errorf("expected (nil, false) from empty pointer queue Peek, got (%v, %v)", val, ok)
	}
}
