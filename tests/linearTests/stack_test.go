package tests

import (
	"testing"

	"datastructures/linear"
)

// --- map type ---

func TestStackWithMaps(t *testing.T) {
	s := linear.StackFixed[map[string]int](3)

	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"b": 2, "c": 3}
	var nilMap map[string]int

	s.Push(m1)
	s.Push(m2)
	s.Push(nilMap)

	if s.Size() != 3 {
		t.Fatalf("expected 3 items, got %d", s.Size())
	}

	top := s.Peek()
	if top != nil {
		t.Errorf("expected nil map on top, got %v", top)
	}

	popped, err := s.Pop()
	if err != nil || popped != nil {
		t.Errorf("expected nil map from Pop, got %v err: %v", popped, err)
	}

	if s.Size() != 2 {
		t.Errorf("expected 2 items after Pop, got %d", s.Size())
	}
}

func TestStackMapMutationAfterPush(t *testing.T) {
	s := linear.StackFixed[map[string]int](2)

	m := map[string]int{"x": 10}
	s.Push(m)

	m["x"] = 999

	top := s.Peek()
	if top["x"] != 999 {
		t.Logf("map mutation not reflected (value copied): got %d", top["x"])
	} else {
		t.Logf("map mutation reflected (reference held): got %d", top["x"])
	}
}

func TestStackMapOverCapacity(t *testing.T) {
	s := linear.StackFixed[map[string]int](1)
	s.Push(map[string]int{"a": 1})

	err := s.Push(map[string]int{"b": 2})
	if err == nil {
		t.Error("expected error pushing map beyond capacity, got nil")
	}
}

// --- struct (object) type ---

func TestStackWithStructs(t *testing.T) {
	s := linear.StackFixed[Person](3)

	s.Push(Person{"Alice", 30})
	s.Push(Person{"Bob", 25})
	s.Push(Person{}) // zero value

	if s.Size() != 3 {
		t.Fatalf("expected 3 items, got %d", s.Size())
	}

	top := s.Peek()
	if top.Name != "" {
		t.Errorf("expected zero-value struct on top, got %v", top)
	}

	bottom := s.First()
	if bottom.Name != "Alice" {
		t.Errorf("expected Alice at bottom, got %v", bottom)
	}

	popped, err := s.Pop()
	if err != nil || popped.Name != "" {
		t.Errorf("expected zero-value struct from Pop, got %v", popped)
	}
}

func TestStackStructOverCapacity(t *testing.T) {
	s := linear.StackFixed[Person](2)
	s.Push(Person{"A", 1})
	s.Push(Person{"B", 2})

	err := s.Push(Person{"C", 3})
	if err == nil {
		t.Error("expected error pushing struct beyond capacity, got nil")
	}
}

func TestStackClearStructs(t *testing.T) {
	s := linear.StackFixed[Person](3)
	s.Push(Person{"Alice", 30})
	s.Push(Person{"Bob", 25})

	s.Clear()

	if s.Size() != 0 {
		t.Errorf("expected empty stack after Clear, got %d items", s.Size())
	}
}

// --- pointer type ---

func TestStackWithPointers(t *testing.T) {
	s := linear.StackFixed[*Person](3)

	p1 := &Person{"Alice", 30}
	p2 := &Person{"Bob", 25}

	s.Push(p1)
	s.Push(p2)
	s.Push(nil)

	if s.Size() != 3 {
		t.Fatalf("expected 3 items, got %d", s.Size())
	}

	top := s.Peek()
	if top != nil {
		t.Errorf("expected nil pointer on top, got %v", top)
	}

	bottom := s.First()
	if bottom != p1 {
		t.Errorf("expected p1 at bottom, got %v", bottom)
	}
}

func TestStackNilPointerPop(t *testing.T) {
	s := linear.StackFixed[*Person](2)
	s.Push(nil)

	got, err := s.Pop()
	if err != nil {
		t.Fatalf("unexpected error popping nil pointer: %v", err)
	}
	if got != nil {
		t.Errorf("expected nil pointer back, got %v", got)
	}
}

func TestStackPointerMutationAfterPush(t *testing.T) {
	s := linear.StackFixed[*Person](1)

	p := &Person{"Alice", 30}
	s.Push(p)

	p.Name = "Mutated"

	top := s.Peek()
	if top.Name != "Mutated" {
		t.Errorf("expected mutation reflected via pointer, got %s", top.Name)
	}
}

// --- slice type ---

func TestStackWithSlices(t *testing.T) {
	s := linear.StackFixed[[]int](3)

	s.Push([]int{1, 2, 3})
	s.Push([]int{})
	s.Push(nil)

	if s.Size() != 3 {
		t.Fatalf("expected 3 items, got %d", s.Size())
	}

	top := s.Peek()
	if top != nil {
		t.Errorf("expected nil slice on top, got %v", top)
	}

	bottom := s.First()
	if len(bottom) != 3 || bottom[0] != 1 {
		t.Errorf("expected [1 2 3] at bottom, got %v", bottom)
	}
}

func TestStackSliceMutationAfterPush(t *testing.T) {
	s := linear.StackFixed[[]int](1)

	sl := []int{1, 2, 3}
	s.Push(sl)

	sl[0] = 999

	top := s.Peek()
	if top[0] != 999 {
		t.Logf("slice mutation not reflected (copied): got %d", top[0])
	} else {
		t.Logf("slice mutation reflected (shared backing array): got %d", top[0])
	}
}

func TestStackSliceOverCapacity(t *testing.T) {
	s := linear.StackFixed[[]int](1)
	s.Push([]int{1})

	err := s.Push([]int{2})
	if err == nil {
		t.Error("expected error pushing slice beyond capacity, got nil")
	}
}

// --- empty stack edge cases ---

func TestStackPopEmpty(t *testing.T) {
	s := linear.StackFixed[map[string]int](2)
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error popping from empty map stack")
	}
}

func TestStackPopEmptyPointer(t *testing.T) {
	s := linear.StackFixed[*Person](2)
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error popping from empty pointer stack")
	}
}

func TestStackPopEmptySlice(t *testing.T) {
	s := linear.StackFixed[[]int](2)
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error popping from empty slice stack")
	}
}

func TestStackPopEmptyStruct(t *testing.T) {
	s := linear.StackFixed[Person](2)
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error popping from empty struct stack")
	}
}

// Peek and First have no bounds check — they panic on empty stacks.
// These tests document that behavior.

func TestStackPeekEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from Peek on empty stack, got none")
		}
	}()
	s := linear.StackFixed[*Person](2)
	s.Peek()
}

func TestStackFirstEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from First on empty stack, got none")
		}
	}()
	s := linear.StackFixed[*Person](2)
	s.First()
}
