package tests

import (
	"slices"
	"testing"

	"datastructures/linear"
)

// --- helpers ---

func intBuffer(capacity int, vals ...int) *linear.CircularBuffer[int] {
	cb := linear.NewCircularBuffer[int](capacity)
	for _, v := range vals {
		cb.Push(v)
	}
	return cb
}

// --- construction ---

func TestCircularBuffer_EmptyData(t *testing.T) {
	cb := linear.NewCircularBuffer[int](10)
	data := cb.Data()
	if len(data) != 10 {
		t.Fatalf("expected backing slice of length 10, got %d", len(data))
	}
	for i, v := range data {
		if v != 0 {
			t.Errorf("expected zero at index %d, got %d", i, v)
		}
	}
}

func TestCircularBuffer_RespectsCapacity(t *testing.T) {
	for _, capacity := range []int{1, 3, 5, 16} {
		cb := linear.NewCircularBuffer[int](capacity)
		if got := len(cb.Data()); got != capacity {
			t.Errorf("capacity %d: expected backing slice length %d, got %d", capacity, capacity, got)
		}
	}
}

// --- basic add ---

func TestCircularBuffer_AddOne(t *testing.T) {
	cb := intBuffer(4, 42)
	if cb.Data()[0] != 42 {
		t.Errorf("expected 42 at index 0, got %d", cb.Data()[0])
	}
}

func TestCircularBuffer_AddMultiple(t *testing.T) {
	cb := intBuffer(5, 1, 2, 3, 4, 5)
	data := cb.Data()
	for i := 0; i < 5; i++ {
		if data[i] != i+1 {
			t.Errorf("index %d: expected %d, got %d", i, i+1, data[i])
		}
	}
}

// --- fill to exact capacity ---

func TestCircularBuffer_FillExact(t *testing.T) {
	cb := intBuffer(5, 1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}
	if got := cb.Data(); !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

// --- overflow: oldest value must be overwritten ---

func TestCircularBuffer_OverflowByOne(t *testing.T) {
	// fill 1..5 into a cap-5 buffer, then add 6; slot 0 gets overwritten.
	cb := intBuffer(5, 1, 2, 3, 4, 5, 6)
	want := []int{6, 2, 3, 4, 5}
	if got := cb.Data(); !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestCircularBuffer_OverflowBy3(t *testing.T) {
	cb := intBuffer(5, 1, 2, 3, 4, 5, 6, 7, 8)
	want := []int{6, 7, 8, 4, 5}
	if got := cb.Data(); !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

// --- double wrap: add exactly 2x capacity ---

func TestCircularBuffer_DoubleWrap(t *testing.T) {
	// add 10 values to a cap-5 buffer; all original slots replaced by 6..10.
	cb := intBuffer(5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	want := []int{6, 7, 8, 9, 10}
	if got := cb.Data(); !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

// --- stress: 100 additions into a cap-10 buffer ---

func TestCircularBuffer_StressAdd(t *testing.T) {
	cb := linear.NewCircularBuffer[int](10)
	for i := 1; i <= 100; i++ {
		cb.Push(i)
	}
	// last 10 values added were 91..100, landing in slots 0..9 in order.
	want := []int{91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	if got := cb.Data(); !slices.Equal(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

// --- string type ---

func TestCircularBuffer_StringType(t *testing.T) {
	cb := linear.NewCircularBuffer[string](4)
	cb.Push("hello")
	cb.Push("world")
	data := cb.Data()
	if data[0] != "hello" {
		t.Errorf("expected hello, got %s", data[0])
	}
	if data[1] != "world" {
		t.Errorf("expected world, got %s", data[1])
	}
}

func TestCircularBuffer_StringOverflow(t *testing.T) {
	cb := linear.NewCircularBuffer[string](5)
	for _, w := range []string{"a", "b", "c", "d", "e", "f"} {
		cb.Push(w)
	}
	// "f" overwrote slot 0 ("a").
	if cb.Data()[0] != "f" {
		t.Errorf("expected 'f' at slot 0, got '%s'", cb.Data()[0])
	}
}

// --- float type ---

func TestCircularBuffer_FloatType(t *testing.T) {
	cb := linear.NewCircularBuffer[float64](4)
	cb.Push(3.14)
	cb.Push(2.71)
	data := cb.Data()
	if data[0] != 3.14 {
		t.Errorf("expected 3.14, got %f", data[0])
	}
	if data[1] != 2.71 {
		t.Errorf("expected 2.71, got %f", data[1])
	}
}

// --- idempotent Data() call ---

func TestCircularBuffer_DataDoesNotMutate(t *testing.T) {
	cb := intBuffer(5, 1, 2, 3)
	first := cb.Data()
	second := cb.Data()
	if !slices.Equal(first, second) {
		t.Errorf("consecutive Data() calls differ: %v vs %v", first, second)
	}
}

// --- untouched slots stay zero ---

func TestCircularBuffer_UnusedSlotsAreZero(t *testing.T) {
	cb := intBuffer(10, 7) // only slot 0 written
	data := cb.Data()
	for i := 1; i < 10; i++ {
		if data[i] != 0 {
			t.Errorf("slot %d: expected 0, got %d", i, data[i])
		}
	}
}
