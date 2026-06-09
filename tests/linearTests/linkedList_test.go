package tests

import (
	"testing"

	"datastructures/linear"
)

// --- InsertAtHead ---

func TestLinkedListInsertAtHeadSingle(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(42)

	got := ll.Data()
	if got == nil {
		t.Fatal("expected non-nil Data after InsertAtHead, got nil")
	}
	if *got != 42 {
		t.Errorf("expected 42, got %d", *got)
	}
}

func TestLinkedListInsertAtHeadOrderIsLIFO(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(1)
	ll.InsertAtHead(2)
	ll.InsertAtHead(3)

	got := ll.Data()
	if got == nil || *got != 3 {
		t.Errorf("expected 3 (last inserted) at head, got %v", got)
	}
}

// --- InsertAtTail ---

func TestLinkedListInsertAtTailSingle(t *testing.T) {
	ll := &linear.LinkedList[string]{}
	ll.InsertAtTail("hello")

	got := ll.Data()
	if got == nil || *got != "hello" {
		t.Errorf("expected 'hello' at head, got %v", got)
	}
}

func TestLinkedListInsertAtTailDoesNotMoveHead(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	got := ll.Data()
	if got == nil || *got != 1 {
		t.Errorf("expected head to remain 1 after tail inserts, got %v", got)
	}
}

// --- Data ---

func TestLinkedListDataOnEmpty(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	if ll.Data() != nil {
		t.Error("expected nil Data on empty list")
	}
}

// --- DeleteHead ---

func TestLinkedListDeleteHeadOnEmpty(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.DeleteHead() // must not panic
	if ll.Data() != nil {
		t.Error("expected nil Data after DeleteHead on empty list")
	}
}

func TestLinkedListDeleteHeadSingleElement(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(5)
	ll.DeleteHead()

	if ll.Data() != nil {
		t.Error("expected nil after deleting the only element via DeleteHead")
	}
}

func TestLinkedListDeleteHeadAdvancesHead(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(1)
	ll.InsertAtHead(2)
	ll.InsertAtHead(3)

	ll.DeleteHead()

	got := ll.Data()
	if got == nil || *got != 2 {
		t.Errorf("expected 2 after deleting head (3), got %v", got)
	}
}

// --- DeleteTail ---

func TestLinkedListDeleteTailOnEmpty(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.DeleteTail() // must not panic
	if ll.Data() != nil {
		t.Error("expected nil Data after DeleteTail on empty list")
	}
}

func TestLinkedListDeleteTailSingleElement(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtTail(10)
	ll.DeleteTail()

	if ll.Data() != nil {
		t.Error("expected nil after deleting the only element via DeleteTail")
	}
}

func TestLinkedListDeleteTailRemovesLastElement(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	ll.DeleteTail()

	// walk head → next to verify 3 is gone
	got := ll.Data()
	if got == nil || *got != 1 {
		t.Errorf("expected head to remain 1, got %v", got)
	}
	ll.DeleteHead()
	got = ll.Data()
	if got == nil || *got != 2 {
		t.Errorf("expected second element to be 2 (3 was deleted), got %v", got)
	}
	ll.DeleteHead()
	if ll.Data() != nil {
		t.Error("expected empty list after removing all elements")
	}
}

// --- string type ---

func TestLinkedListWithStrings(t *testing.T) {
	ll := &linear.LinkedList[string]{}
	ll.InsertAtHead("bob")
	ll.InsertAtHead("alice")

	got := ll.Data()
	if got == nil || *got != "alice" {
		t.Errorf("expected alice at head, got %v", got)
	}

	ll.DeleteHead()
	got = ll.Data()
	if got == nil || *got != "bob" {
		t.Errorf("expected bob after deleting alice, got %v", got)
	}
}

func TestLinkedListZeroValue(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(0)

	got := ll.Data()
	if got == nil || *got != 0 {
		t.Errorf("expected zero-value int at head, got %v", got)
	}
}

// --- pointer type ---

func TestLinkedListWithPointers(t *testing.T) {
	ll := &linear.LinkedList[*int]{}
	v1, v2 := 10, 20
	p1, p2 := &v1, &v2

	ll.InsertAtTail(p1)
	ll.InsertAtTail(p2)

	got := ll.Data()
	if got == nil || *got != p1 {
		t.Errorf("expected p1 at head, got %v", got)
	}
}

func TestLinkedListPointerMutation(t *testing.T) {
	ll := &linear.LinkedList[*int]{}
	v := 42
	p := &v
	ll.InsertAtHead(p)

	v = 99

	got := ll.Data()
	if got == nil || **got != 99 {
		t.Errorf("expected mutation reflected via pointer, got %v", got)
	}
}

func TestLinkedListNilPointer(t *testing.T) {
	ll := &linear.LinkedList[*int]{}
	ll.InsertAtHead(nil)

	got := ll.Data()
	if got == nil {
		t.Fatal("expected non-nil Data pointer (pointing to a nil *int), got nil")
	}
	if *got != nil {
		t.Errorf("expected stored nil *int, got %v", *got)
	}
}

// --- mixed operations ---

func TestLinkedListAlternatingInsertDelete(t *testing.T) {
	ll := &linear.LinkedList[int]{}
	ll.InsertAtHead(1)
	ll.InsertAtTail(2)
	ll.InsertAtHead(0)
	// list: 0 <--> 1 <--> 2

	ll.DeleteTail()
	// list: 0 <--> 1

	got := ll.Data()
	if got == nil || *got != 0 {
		t.Errorf("expected 0 at head, got %v", got)
	}

	ll.DeleteHead()
	// list: 1
	got = ll.Data()
	if got == nil || *got != 1 {
		t.Errorf("expected 1 as sole remaining element, got %v", got)
	}

	ll.DeleteHead()
	if ll.Data() != nil {
		t.Error("expected empty list after deleting all elements")
	}
}

func TestLinkedListDeleteBothEndsToEmpty(t *testing.T) {
	ll := &linear.LinkedList[string]{}
	ll.InsertAtTail("a")
	ll.InsertAtTail("b")

	ll.DeleteHead()
	ll.DeleteTail()

	if ll.Data() != nil {
		t.Error("expected empty list after removing both elements")
	}
}
