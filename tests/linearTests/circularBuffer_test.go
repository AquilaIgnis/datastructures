package tests

import (
	"testing"

	"datastructures/linear"
)

// --- helpers ---

func intBuffer(vals ...int) linear.CircularBuffer[int] {
	var cb linear.CircularBuffer[int]
	for _, v := range vals {
		cb.Push(v)
	}
	return cb
}

// --- empty buffer ---

func TestCircularBuffer_EmptyData(t *testing.T) {
	var cb linear.CircularBuffer[int]
	data := cb.Data()
	for i, v := range data {
		if v != 0 {
			t.Errorf("expected zero at index %d, got %d", i, v)
		}
	}
}

// --- basic add ---

func TestCircularBuffer_AddOne(t *testing.T) {
	cb := intBuffer(42)
	if cb.Data()[0] != 42 {
		t.Errorf("expected 42 at index 0, got %d", cb.Data()[0])
	}
}

func TestCircularBuffer_AddMultiple(t *testing.T) {
	cb := intBuffer(1, 2, 3, 4, 5)
	data := cb.Data()
	for i := 0; i < 5; i++ {
		if data[i] != i+1 {
			t.Errorf("index %d: expected %d, got %d", i, i+1, data[i])
		}
	}
}

// --- fill to exact capacity (10) ---

func TestCircularBuffer_FillExact(t *testing.T) {
	cb := intBuffer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	data := cb.Data()
	for i := 0; i < 10; i++ {
		if data[i] != i+1 {
			t.Errorf("index %d: expected %d, got %d", i, i+1, data[i])
		}
	}
}

// --- overflow: oldest value must be overwritten ---

func TestCircularBuffer_OverflowByOne(t *testing.T) {
	// fill with 1..10 then add 11; slot 0 gets overwritten
	cb := intBuffer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	data := cb.Data()
	if data[0] != 11 {
		t.Errorf("expected slot 0 to be overwritten with 11, got %d", data[0])
	}
	// slots 1-9 should still hold 2-10
	for i := 1; i < 10; i++ {
		if data[i] != i+1 {
			t.Errorf("slot %d: expected %d, got %d", i, i+1, data[i])
		}
	}
}

func TestCircularBuffer_OverflowBy5(t *testing.T) {
	// add 15 values: last 10 should be 6..15
	cb := intBuffer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	data := cb.Data()
	// slots 0-4 are overwritten with 11-15
	expected := [10]int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10}
	for i, want := range expected {
		if data[i] != want {
			t.Errorf("slot %d: expected %d, got %d", i, want, data[i])
		}
	}
}

// --- double wrap: add exactly 2x capacity ---

func TestCircularBuffer_DoubleWrap(t *testing.T) {
	// add 20 values; all original slots fully replaced by 11-20
	cb := intBuffer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	data := cb.Data()
	expected := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, want := range expected {
		if data[i] != want {
			t.Errorf("slot %d: expected %d, got %d", i, want, data[i])
		}
	}
}

// --- stress: 100 additions ---

func TestCircularBuffer_StressAdd(t *testing.T) {
	var cb linear.CircularBuffer[int]
	for i := 1; i <= 100; i++ {
		cb.Push(i)
	}
	// last 10 values added were 91-100
	// after 100 adds: tail wraps, slots should hold 91-100
	data := cb.Data()
	// slot index = (i-1) % 10, value = i for i in 91..100
	expected := [10]int{91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	for i, want := range expected {
		if data[i] != want {
			t.Errorf("slot %d: expected %d, got %d", i, want, data[i])
		}
	}
}

// --- string type ---

func TestCircularBuffer_StringType(t *testing.T) {
	var cb linear.CircularBuffer[string]
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
	var cb linear.CircularBuffer[string]
	words := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	for _, w := range words {
		cb.Push(w)
	}
	// "k" overwrote slot 0 ("a")
	if cb.Data()[0] != "k" {
		t.Errorf("expected 'k' at slot 0, got '%s'", cb.Data()[0])
	}
}

// --- float type ---

func TestCircularBuffer_FloatType(t *testing.T) {
	var cb linear.CircularBuffer[float64]
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
	cb := intBuffer(1, 2, 3)
	first := cb.Data()
	second := cb.Data()
	if first != second {
		t.Error("consecutive Data() calls returned different results")
	}
}

// --- zero value after overflow stays zero for untouched slots ---

func TestCircularBuffer_UnusedSlotsAreZero(t *testing.T) {
	cb := intBuffer(7) // only slot 0 written
	data := cb.Data()
	for i := 1; i < 10; i++ {
		if data[i] != 0 {
			t.Errorf("slot %d: expected 0, got %d", i, data[i])
		}
	}
}
