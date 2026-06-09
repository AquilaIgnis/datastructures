package tests

import (
	"testing"

	"datastructures/algo"
)

// --- basic profit cases ---

func TestMaxSubarray_BasicProfit(t *testing.T) {
	// buy at 1, sell at 6 → 5
	got := algo.MaxSubarray([]int{7, 1, 5, 3, 6, 4})
	if got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
}

func TestMaxSubarray_ProfitAtEnd(t *testing.T) {
	// buy at 1, sell at 9 → 8
	got := algo.MaxSubarray([]int{3, 1, 4, 1, 5, 9})
	if got != 8 {
		t.Errorf("expected 8, got %d", got)
	}
}

func TestMaxSubarray_AlreadySorted(t *testing.T) {
	// buy at 1, sell at 5 → 4
	got := algo.MaxSubarray([]int{1, 2, 3, 4, 5})
	if got != 4 {
		t.Errorf("expected 4, got %d", got)
	}
}

// --- no profit cases ---

func TestMaxSubarray_StrictlyDecreasing(t *testing.T) {
	// prices only fall; no profitable trade
	got := algo.MaxSubarray([]int{5, 4, 3, 2, 1})
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestMaxSubarray_AllSame(t *testing.T) {
	got := algo.MaxSubarray([]int{3, 3, 3, 3})
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

// --- single element ---

func TestMaxSubarray_SingleElement(t *testing.T) {
	got := algo.MaxSubarray([]int{42})
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

// --- two elements ---

func TestMaxSubarray_TwoElementsProfit(t *testing.T) {
	got := algo.MaxSubarray([]int{1, 10})
	if got != 9 {
		t.Errorf("expected 9, got %d", got)
	}
}

func TestMaxSubarray_TwoElementsLoss(t *testing.T) {
	got := algo.MaxSubarray([]int{10, 1})
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

// --- minimum appears late ---

func TestMaxSubarray_MinInMiddle(t *testing.T) {
	// dip to 2 at index 3, then rises to 8
	got := algo.MaxSubarray([]int{5, 4, 3, 2, 8})
	if got != 6 {
		t.Errorf("expected 6, got %d", got)
	}
}

// --- negative values ---

func TestMaxSubarray_AllNegative(t *testing.T) {
	// minimum tracks most negative; no positive spread possible
	got := algo.MaxSubarray([]int{-5, -3, -1})
	if got != 4 {
		t.Errorf("expected 4, got %d", got)
	}
}

func TestMaxSubarray_MixedNegativePositive(t *testing.T) {
	// buy at -3, sell at 5 → 8
	got := algo.MaxSubarray([]int{2, -3, 1, 5})
	if got != 8 {
		t.Errorf("expected 8, got %d", got)
	}
}
