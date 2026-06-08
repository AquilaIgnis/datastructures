package algo

// kadane's algorithm
// https://www.youtube.com/watch?v=qj3CjNEKFeM

// MaxSubarray  -> returns the max sum of possible sub arrays
// O(n)
func MaxSubarray(array []int) int {
	minimum := array[0]
	maximum := 0

	for i := 1; i < len(array); i++ {
		if array[i] < minimum {
			minimum = array[i]
		} else if array[i]-minimum > maximum {
			maximum = array[i] - minimum
		}
	}

	return maximum
}
