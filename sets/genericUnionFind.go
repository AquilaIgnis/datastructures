package sets

import "cmp"

// GenericUnionFind{
// unionFind *UnionFind
// index     map[T]int
// names     []T
// }
// A union-find over any cmp.Ordered type, backed by an integer UnionFind.
type GenericUnionFind[T cmp.Ordered] struct {
	unionFind *UnionFind
	index     map[T]int
	names     []T
}

// NewGenericUnion() -> creates a disjoint set from (generic comparable)
// the constructor will deduplicate and return a GenericUnionFind
func NewGenericUnion[T cmp.Ordered](params []T) *GenericUnionFind[T] {
	genericUnion := &GenericUnionFind[T]{
		index: make(map[T]int, len(params)),
	}

	for _, name := range params {
		if _, seen := genericUnion.index[name]; !seen {
			genericUnion.index[name] = len(genericUnion.names)
			genericUnion.names = append(genericUnion.names, name)
		}
	}

	genericUnion.unionFind = NewUnionFind(len(genericUnion.names))

	return genericUnion
}

// Union() -> creates an Union of 2 values, returns true if success
func (n *GenericUnionFind[T]) Union(a, b T) bool {
	paramA, ok1 := n.index[a]
	paramB, ok2 := n.index[b]
	if !ok1 || !ok2 {
		return false
	}
	return n.unionFind.Union(paramA, paramB)
}

// IsUnion() -> returns true if the parameters have the same root
func (n *GenericUnionFind[T]) IsUnion(a T, b T) bool {
	paramA, ok1 := n.index[a]
	paramB, ok2 := n.index[b]
	if !ok1 || !ok2 {
		return false
	}
	return n.unionFind.IsUnion(paramA, paramB)
}

// Rep() -> returns the representative of  group, and false if the root is unknown.
func (n *GenericUnionFind[T]) Rep(a T) (T, bool) {
	ia, ok := n.index[a]
	if !ok {
		var zero T
		return zero, false
	}
	return n.names[n.unionFind.Find(ia)], true
}

// Disjointed() -> returns the count of disjointed sets remaining
func (n *GenericUnionFind[T]) Disjointed() int {
	return n.unionFind.Disjointed()
}
