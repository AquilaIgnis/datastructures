package tests

import (
	"testing"

	"datastructures/linear"
)

// --- Current on an untouched buffer ---

func TestCircularBuffer_CurrentNilWhenEmpty(t *testing.T) {
	var cb linear.CircularBuffer[int]
	if cb.Current != nil {
		t.Errorf("expected Current to be nil on empty buffer, got %v", *cb.Current)
	}
}

func TestCircularBuffer_CurrentNilFromConstructor(t *testing.T) {
	cb := linear.NewCircularBuffer[int]()
	if cb.Current != nil {
		t.Errorf("expected Current to be nil from NewCircularBuffer, got %v", *cb.Current)
	}
}

// --- Current points at the most recently pushed value ---

func TestCircularBuffer_CurrentAfterOnePush(t *testing.T) {
	var cb linear.CircularBuffer[int]
	cb.Push(42)
	if cb.Current == nil {
		t.Fatal("expected Current to be set after Push, got nil")
	}
	if *cb.Current != 42 {
		t.Errorf("expected *Current == 42, got %d", *cb.Current)
	}
}

func TestCircularBuffer_CurrentTracksLatest(t *testing.T) {
	var cb linear.CircularBuffer[int]
	for _, v := range []int{1, 2, 3, 4, 5} {
		cb.Push(v)
		if *cb.Current != v {
			t.Errorf("after pushing %d, expected *Current == %d, got %d", v, v, *cb.Current)
		}
	}
}

// --- Current aliases the live backing slot ---

func TestCircularBuffer_CurrentAliasesWrittenSlot(t *testing.T) {
	var cb linear.CircularBuffer[int]
	cb.Push(10)
	cb.Push(20)
	cb.Push(30)
	// Last write landed in slot 2; Data() is a snapshot of the backing array.
	if got := cb.Data()[2]; *cb.Current != got {
		t.Errorf("*Current (%d) should match Data()[2] (%d)", *cb.Current, got)
	}
}

// --- Current follows the write head through a wrap-around ---

func TestCircularBuffer_CurrentAfterOverflow(t *testing.T) {
	var cb linear.CircularBuffer[int]
	for i := 1; i <= 11; i++ {
		cb.Push(i)
	}
	// The 11th value (11) overwrote slot 0 and is the most recent write.
	if *cb.Current != 11 {
		t.Errorf("expected *Current == 11 after overflow, got %d", *cb.Current)
	}
	if got := cb.Data()[0]; *cb.Current != got {
		t.Errorf("*Current (%d) should match overwritten slot 0 (%d)", *cb.Current, got)
	}
}

// --- Current works for non-int types ---

func TestCircularBuffer_CurrentStringType(t *testing.T) {
	var cb linear.CircularBuffer[string]
	cb.Push("hello")
	cb.Push("world")
	if cb.Current == nil || *cb.Current != "world" {
		t.Errorf("expected *Current == \"world\", got %v", cb.Current)
	}
}
