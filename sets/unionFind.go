package sets

//	UnionFind {
//		disjointed int
//		parent     []int
//		rank       []int
//	}
type UnionFind struct {
	disjointed int
	parent     []int
	rank       []int
}

// NewUnionFind() -> Creates an disjoint set ,takes the unique elements as parameter.
func NewUnionFind(uniqueElements int) *UnionFind {
	// counts from 0 correction
	uniqueElements += 1

	uf := &UnionFind{
		parent:     make([]int, uniqueElements),
		rank:       make([]int, uniqueElements),
		disjointed: uniqueElements - 1,
	}

	for i := range uf.parent {
		uf.parent[i] = i
	}

	return uf
}

// Find() -> returns the root of a node
func (uf *UnionFind) Find(val int) int {
	root := val

	// find root of val
	for uf.parent[root] != root {
		root = uf.parent[root]
	}

	// point every node on the path directly at root
	current := val
	for uf.parent[current] != root {
		current, uf.parent[current] = uf.parent[current], root
	}

	return root
}

// Union() -> creates an Union of 2 values, returns true if success
func (uf *UnionFind) Union(a int, b int) bool {
	node1 := uf.Find(a)
	node2 := uf.Find(b)

	if node1 == node2 {
		return false
	}

	winner, loser := node1, node2
	if uf.rank[winner] < uf.rank[loser] {
		winner, loser = loser, winner
	}

	// merge
	uf.parent[loser] = winner

	if uf.rank[winner] == uf.rank[loser] {
		uf.rank[winner]++
	}
	uf.disjointed--

	return true
}

// IsUnion() -> returns true if the parameters have the same root
func (uf *UnionFind) IsUnion(a int, b int) bool {
	if uf.Find(a) == uf.Find(b) {
		return true
	}
	return false
}

// Disjointed() -> returns the count of disjointed sets remaining
func (uf *UnionFind) Disjointed() int {
	return uf.disjointed
}
