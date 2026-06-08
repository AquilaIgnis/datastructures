package algo

// Fib() -> fibonacci linear implementation O(n)
func Fib(target int) int {
	a := 0
	b := 1

	if target <= 1 {
		return target
	}

	for a < target {
		a, b = b, a+b
	}
	return a
}
