package linear

// BaseTypes
type BaseTypes interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string | bool
}

// CircularBuffer -> Circular Buffer struct only takes baseTypes
// .Current is a pointer to tail
type CircularBuffer[T BaseTypes] struct {
	array   []T
	cHead   int
	cTail   int
	size    int
	Current *T
}

// NewCircularBuffer[type](size) -> Creates a circular buffer of given size
func NewCircularBuffer[T BaseTypes](size int) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		array: make([]T, size),
	}
}

// Push() -> adds values wrapping around to overwrite oldest
func (cb *CircularBuffer[T]) Push(val T) {
	cb.array[cb.cTail] = val
	cb.Current = &cb.array[cb.cTail]

	cb.cTail = (cb.cTail + 1) % len(cb.array) // wrap around
	if cb.size < len(cb.array) {
		cb.size++
	} else {
		cb.cHead = (cb.cHead + 1) % len(cb.array) // overwrite oldest
	}
}

// Data() -> returns the entire array in CircularBuffer
func (cb CircularBuffer[T]) Data() []T {
	return cb.array
}
