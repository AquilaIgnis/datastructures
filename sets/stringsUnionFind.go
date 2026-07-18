package sets

// NamedUnionFind{
// unionFind *UnionFind
// index     map[string]int
// names     []string
// }
type NamedUnionFind struct {
	unionFind *UnionFind
	index     map[string]int
	names     []string
}

// NewStringsUnion() -> creates a disjoint set from a slice of strings,
// the constructor will deduplicate and return a NamedUnionFind
func NewStringsUnion(params []string) *NamedUnionFind {
	namedUnion := &NamedUnionFind{
		index: make(map[string]int, len(params)),
	}

	for _, name := range params {
		if _, seen := namedUnion.index[name]; !seen {
			namedUnion.index[name] = len(namedUnion.names)
			namedUnion.names = append(namedUnion.names, name)
		}
	}

	namedUnion.unionFind = NewUnionFind(len(namedUnion.names))

	return namedUnion
}

// Union() -> creates an Union of 2 values, returns true if success
func (n *NamedUnionFind) Union(a, b string) bool {
	paramA, ok1 := n.index[a]
	paramB, ok2 := n.index[b]
	if !ok1 || !ok2 {
		return false
	}
	return n.unionFind.Union(paramA, paramB)
}

// IsUnion() -> returns true if the parameters have the same root
func (n *NamedUnionFind) IsUnion(a, b string) bool {
	paramA, ok1 := n.index[a]
	paramB, ok2 := n.index[b]
	if !ok1 || !ok2 {
		return false
	}
	return n.unionFind.IsUnion(paramA, paramB)
}

// Rep() -> returns the representative name of  group, and false if the name is unknown.
func (n *NamedUnionFind) Rep(a string) (string, bool) {
	ia, ok := n.index[a]
	if !ok {
		return "", false
	}
	return n.names[n.unionFind.Find(ia)], true
}

// Disjointed() -> returns the count of disjointed sets remaining
func (n *NamedUnionFind) Disjointed() int {
	return n.unionFind.Disjointed()
}
