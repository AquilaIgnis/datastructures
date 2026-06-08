package algo

// GCD  gives the greatest common factor of two numbers
// EuclideanGCD  O(log(n))
func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
