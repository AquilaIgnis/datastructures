package tests

import (
	"testing"

	"datastructures/algo"
)

// algo.MaxSubarray implements Kadane-style "max profit": it returns the largest
// value of array[j] - array[i] for i < j, or 0 when no such positive difference
// exists. (It is a buy-low / sell-high spread, not a contiguous-sum.)

func TestMaxSubarray(t *testing.T) {
	cases := []struct {
		name  string
		array []int
		want  int
	}{
		{"basic profit", []int{7, 1, 5, 3, 6, 4}, 5},
		{"profit at end", []int{3, 1, 4, 1, 5, 9}, 8},
		{"already sorted", []int{1, 2, 3, 4, 5}, 4},
		{"strictly decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"all same", []int{3, 3, 3, 3}, 0},
		{"single element", []int{42}, 0},
		{"two elements profit", []int{1, 10}, 9},
		{"two elements loss", []int{10, 1}, 0},
		{"min in middle then rises", []int{5, 4, 3, 2, 8}, 6},
		{"all negative", []int{-5, -3, -1}, 4},
		{"mixed neg/pos", []int{2, -3, 1, 5}, 8},
		{"dip late", []int{9, 8, 7, 1, 2}, 1},
		{"zero crossing", []int{-1, 0, -2, 3}, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := algo.MaxSubarray(c.array); got != c.want {
				t.Errorf("MaxSubarray(%v) = %d, want %d", c.array, got, c.want)
			}
		})
	}
}

// bruteForceMaxProfit is the obvious O(n^2) reference: best later-minus-earlier
// spread, floored at 0.
func bruteForceMaxProfit(array []int) int {
	best := 0
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if d := array[j] - array[i]; d > best {
				best = d
			}
		}
	}
	return best
}

// Property: the linear implementation must agree with the brute-force reference
// across a deterministic spread of inputs.
func TestMaxSubarray_MatchesBruteForce(t *testing.T) {
	inputs := [][]int{
		{1},
		{2, 1},
		{1, 2},
		{4, 1, 7, 2, 9, 3},
		{-2, -5, -1, -8, -3},
		{0, 0, 0, 1},
		{10, -10, 10, -10, 20},
		{3, 3, 4, 2, 5, 1, 6},
		{-1, 2, -3, 4, -5, 6},
		{100, 50, 75, 25, 80},
	}

	for _, in := range inputs {
		want := bruteForceMaxProfit(in)
		if got := algo.MaxSubarray(in); got != want {
			t.Errorf("MaxSubarray(%v) = %d, brute force = %d", in, got, want)
		}
	}
}
