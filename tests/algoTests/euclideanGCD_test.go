package tests

import (
	"testing"

	"datastructures/algo"
)

func TestGCD(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"coprime", 13, 7, 1},
		{"common factor", 12, 8, 4},
		{"one divides other", 21, 7, 7},
		{"equal values", 9, 9, 9},
		{"b is zero", 5, 0, 5},
		{"a is zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
		{"a less than b", 8, 12, 4},
		{"one", 1, 999, 1},
		{"large coprime primes", 17 * 19, 23, 1},
		{"large common factor", 1071, 462, 21},
		{"powers of two", 1024, 256, 256},
		{"consecutive integers are coprime", 100, 101, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := algo.GCD(c.a, c.b); got != c.want {
				t.Errorf("GCD(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
			}
		})
	}
}

// GCD(a, b) should equal GCD(b, a) for the inputs we support.
func TestGCD_Commutative(t *testing.T) {
	pairs := [][2]int{{12, 8}, {21, 7}, {1071, 462}, {13, 7}, {100, 101}}
	for _, p := range pairs {
		if got, rev := algo.GCD(p[0], p[1]), algo.GCD(p[1], p[0]); got != rev {
			t.Errorf("GCD not commutative for (%d, %d): %d != %d", p[0], p[1], got, rev)
		}
	}
}

// The result must evenly divide both inputs (when non-zero).
func TestGCD_DividesBothInputs(t *testing.T) {
	pairs := [][2]int{{12, 8}, {1071, 462}, {1024, 256}, {18, 24}}
	for _, p := range pairs {
		g := algo.GCD(p[0], p[1])
		if g == 0 {
			t.Fatalf("GCD(%d, %d) returned 0 for non-zero inputs", p[0], p[1])
		}
		if p[0]%g != 0 || p[1]%g != 0 {
			t.Errorf("GCD(%d, %d) = %d does not divide both inputs", p[0], p[1], g)
		}
	}
}
