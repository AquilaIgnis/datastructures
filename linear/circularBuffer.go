package linear

type baseTypes interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string | bool
}

// CircularBuffer -> Circular Buffer struct only takes baseTypes
type CircularBuffer[T baseTypes] struct {
	array [10]T
	cHead int
	cTail int
	size  int
}

// Add() -> adds values wrapping around to overwrite oldest
func (cb *CircularBuffer[T]) Add(val T) {
	cb.array[cb.cTail] = val

	cb.cTail = (cb.cTail + 1) % len(cb.array) // wrap around
	if cb.size < len(cb.array) {
		cb.size++
	} else {
		cb.cHead = (cb.cHead + 1) % len(cb.array) // overwrite oldest
	}
}

// Data() -> returns the array in CircularBuffer
func (cb CircularBuffer[T]) Data() [10]T {
	return cb.array
}
