package algo

// Deterministic Miller Rabin primality test

import (
	"math/bits"
)

// mulMod() :: using bits to prevent 2^x overflow
func mulMod(a, b, m uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	_, rem := bits.Div64(hi, lo, m)
	return rem
}

// (base ^ exp) % m  square and multiply
func powMod(base, exp, m uint64) uint64 {
	result := uint64(1)
	base %= m
	for exp > 0 {
		// exp is odd
		if exp&1 == 1 {
			result = mulMod(result, base, m)
		}
		base = mulMod(base, base, m)
		// exp = exp / 2
		exp >>= 1
	}
	return result
}

// IsPrime() -> uses Miller Rabin primality test
func IsPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	// quick trial small primes bases
	smallPrimes := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, p := range smallPrimes {
		if n%p == 0 {
			return n == p
		}
	}

	// n-1 = 2^r * d
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}

	// bases test
	for _, a := range smallPrimes {
		x := powMod(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for i := 0; i < r-1; i++ {
			x = mulMod(x, x, n)
			if x == n-1 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}
