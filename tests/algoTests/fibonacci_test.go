package tests

import (
	"testing"

	"datastructures/algo"
)

// NOTE on behavior:
// algo.Fib does NOT return the target-th Fibonacci number. For target > 1 it
// advances through the Fibonacci sequence (0,1,1,2,3,5,8,13,...) and returns
// the smallest Fibonacci number that is >= target. For target <= 1 it returns
// target unchanged. These tests pin that actual behavior.
//
//   Fib(4) = 5  (smallest fib >= 4), not 3 (the 4th fib)
//   Fib(7) = 8  (smallest fib >= 7), not 13
//
// If the intent was "n-th Fibonacci number", the implementation is buggy and
// these expectations should be updated alongside the fix.

func TestFib(t *testing.T) {
	cases := []struct {
		target int
		want   int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 5},
		{6, 8},
		{7, 8},
		{8, 8},
		{9, 13},
		{12, 13},
		{13, 13},
		{14, 21},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if got := algo.Fib(c.target); got != c.want {
				t.Errorf("Fib(%d) = %d, want %d", c.target, got, c.want)
			}
		})
	}
}

// target <= 1 is returned unchanged, including negative inputs.
func TestFib_SmallAndNegative(t *testing.T) {
	cases := []struct {
		target, want int
	}{
		{-5, -5},
		{-1, -1},
		{0, 0},
		{1, 1},
	}
	for _, c := range cases {
		if got := algo.Fib(c.target); got != c.want {
			t.Errorf("Fib(%d) = %d, want %d", c.target, got, c.want)
		}
	}
}

// Property: for target >= 2 the result is a real Fibonacci number and is the
// smallest such number that is >= target.
func TestFib_ReturnsSmallestFibAtOrAboveTarget(t *testing.T) {
	// Build the set of Fibonacci numbers up to a comfortable bound.
	fibs := []int{0, 1}
	for fibs[len(fibs)-1] < 100000 {
		n := len(fibs)
		fibs = append(fibs, fibs[n-1]+fibs[n-2])
	}
	isFib := func(v int) bool {
		for _, f := range fibs {
			if f == v {
				return true
			}
		}
		return false
	}
	smallestFibGE := func(target int) int {
		for _, f := range fibs {
			if f >= target {
				return f
			}
		}
		t.Fatalf("target %d exceeds precomputed fib range", target)
		return -1
	}

	for target := 2; target <= 1000; target++ {
		got := algo.Fib(target)
		if !isFib(got) {
			t.Fatalf("Fib(%d) = %d is not a Fibonacci number", target, got)
		}
		if got < target {
			t.Fatalf("Fib(%d) = %d is below target", target, got)
		}
		if want := smallestFibGE(target); got != want {
			t.Errorf("Fib(%d) = %d, want smallest fib >= target = %d", target, got, want)
		}
	}
}
