package tests

import (
	"testing"

	"datastructures/linear"
)

// --- map type ---

func TestQueueWithMaps(t *testing.T) {
	q := linear.QueueFixed[map[string]int](3)

	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"b": 2, "c": 3}
	var nilMap map[string]int

	q.Add(m1)
	q.Add(m2)
	q.Add(nilMap)

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	got, ok := q.Peek()
	if !ok || got["a"] != 1 {
		t.Errorf("Peek returned wrong map: %v", got)
	}

	pulled, err := q.Pull()
	if err != nil || pulled["a"] != 1 {
		t.Errorf("Pull returned wrong map: %v err: %v", pulled, err)
	}

	if q.Size() != 2 {
		t.Errorf("expected size 2 after Pull, got %d", q.Size())
	}
}

func TestQueueMapMutationAfterEnqueue(t *testing.T) {
	q := linear.QueueFixed[map[string]int](2)

	m := map[string]int{"x": 10}
	q.Add(m)

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
	q := linear.QueueFixed[map[string]int](1)
	q.Add(map[string]int{"a": 1})

	err := q.Add(map[string]int{"b": 2})
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
	q := linear.QueueFixed[Person](3)

	q.Add(Person{"Alice", 30})
	q.Add(Person{"Bob", 25})
	q.Add(Person{}) // zero value struct

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	first, err := q.Pull()
	if err != nil || first.Name != "Alice" {
		t.Errorf("expected Alice, got %v", first)
	}

	last, err := q.Cull()
	if err != nil || last.Name != "" {
		t.Errorf("expected zero-value struct at tail, got %v", last)
	}
}

func TestQueueStructOverCapacity(t *testing.T) {
	q := linear.QueueFixed[Person](2)
	q.Add(Person{"A", 1})
	q.Add(Person{"B", 2})

	err := q.Add(Person{"C", 3})
	if err == nil {
		t.Error("expected error adding struct beyond capacity, got nil")
	}
}

// --- pointer type ---

func TestQueueWithPointers(t *testing.T) {
	q := linear.QueueFixed[*Person](3)

	p1 := &Person{"Alice", 30}
	p2 := &Person{"Bob", 25}

	q.Add(p1)
	q.Add(p2)
	q.Add(nil) // nil pointer

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	got, ok := q.Peek()
	if !ok || got != p1 {
		t.Errorf("Peek: expected p1 pointer, got %v", got)
	}
}

func TestQueueNilPointerPull(t *testing.T) {
	q := linear.QueueFixed[*Person](2)
	q.Add(nil)

	got, err := q.Pull()
	if err != nil {
		t.Fatalf("unexpected error pulling nil pointer: %v", err)
	}
	if got != nil {
		t.Errorf("expected nil pointer back, got %v", got)
	}
}

func TestQueuePointerMutationAfterEnqueue(t *testing.T) {
	q := linear.QueueFixed[*Person](1)

	p := &Person{"Alice", 30}
	q.Add(p)

	p.Name = "Mutated"

	peeked, _ := q.Peek()
	if peeked.Name != "Mutated" {
		t.Errorf("expected mutation to be reflected via pointer, got %s", peeked.Name)
	}
}

// --- slice type ---

func TestQueueWithSlices(t *testing.T) {
	q := linear.QueueFixed[[]int](3)

	q.Add([]int{1, 2, 3})
	q.Add([]int{}) // empty slice
	q.Add(nil)     // nil slice

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	first, err := q.Pull()
	if err != nil || len(first) != 3 || first[0] != 1 {
		t.Errorf("unexpected first slice: %v", first)
	}
}

func TestQueueSliceMutationAfterEnqueue(t *testing.T) {
	q := linear.QueueFixed[[]int](1)

	s := []int{1, 2, 3}
	q.Add(s)

	s[0] = 999

	peeked, _ := q.Peek()
	if peeked[0] != 999 {
		t.Logf("slice mutation not reflected (copied): got %d", peeked[0])
	} else {
		t.Logf("slice mutation reflected (shared backing array): got %d", peeked[0])
	}
}

func TestQueueSliceOverCapacity(t *testing.T) {
	q := linear.QueueFixed[[]int](1)
	q.Add([]int{1})

	err := q.Add([]int{2})
	if err == nil {
		t.Error("expected error adding slice beyond capacity, got nil")
	}
}

// --- empty queue edge cases across all types ---

func TestQueuePullEmptyMap(t *testing.T) {
	q := linear.QueueFixed[map[string]int](2)
	_, err := q.Pull()
	if err == nil {
		t.Error("expected error pulling from empty map queue")
	}
}

func TestQueuePullEmptyPointer(t *testing.T) {
	q := linear.QueueFixed[*Person](2)
	_, err := q.Pull()
	if err == nil {
		t.Error("expected error pulling from empty pointer queue")
	}
}

func TestQueuePullEmptySlice(t *testing.T) {
	q := linear.QueueFixed[[]int](2)
	_, err := q.Pull()
	if err == nil {
		t.Error("expected error pulling from empty slice queue")
	}
}

func TestQueueCullEmptyStruct(t *testing.T) {
	q := linear.QueueFixed[Person](2)
	_, err := q.Cull()
	if err == nil {
		t.Error("expected error culling from empty struct queue")
	}
}

func TestQueuePeekEmptyPointer(t *testing.T) {
	q := linear.QueueFixed[*Person](2)
	val, ok := q.Peek()
	if ok || val != nil {
		t.Errorf("expected (nil, false) from empty pointer queue Peek, got (%v, %v)", val, ok)
	}
}
