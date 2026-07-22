package graphs

import "slices"

// undirected, unweighted Graph
type Graph[T comparable] struct {
	adjList map[T][]T
}

// NewGraph() -> Creates an unweighted, undirected graph
func NewGraph[T comparable]() *Graph[T] {
	gp := &Graph[T]{
		adjList: make(map[T][]T),
	}

	return gp
}

// AddVertex() -> Creates a new Vertex connecting it to an existing one. Creates them both if
// they dont exist , if both exist then it adds an edge between them.
func (gp *Graph[T]) AddVertex(newVertex T, connectToVertex T) {
	// self loops are not allowed in this graph
	if newVertex == connectToVertex {
		return
	}

	if !slices.Contains(gp.adjList[newVertex], connectToVertex) {
		gp.adjList[newVertex] = append(gp.adjList[newVertex], connectToVertex)
	}

	if !slices.Contains(gp.adjList[connectToVertex], newVertex) {
		gp.adjList[connectToVertex] = append(gp.adjList[connectToVertex], newVertex)
	}
}

// AddIsolatedVertex() -> Creates an isolated Vertex returns True if ok ,
// false if vertex already exist
func (gp *Graph[T]) AddIsolatedVertex(newVertex T) bool {
	_, ok := gp.adjList[newVertex]
	if !ok {
		gp.adjList[newVertex] = make([]T, 0)
		return true
	}
	return false
}

// ConnectEdges() -> Connects an existing vertex to several vertices
func (gp *Graph[T]) ConnectEdges(vertex T, uwu ...T) bool {
	_, ok := gp.adjList[vertex]

	for _, edge := range uwu {

		_, ok = gp.adjList[edge]
		if !ok {
			return ok
		}
	}

	for _, edge := range uwu {
		gp.AddVertex(vertex, edge)
	}

	return ok
}

// Order() -> returns the amount of Vertices
func (gp *Graph[T]) Order() int {
	return len(gp.adjList)
}

// Size() -> returns the amount of Edges
func (gp *Graph[T]) Size() int {
	var size int

	for _, val := range gp.adjList {
		size += len(val)
	}
	return size / 2
}

// Display() -> returns adjList[map]T []T
func (gp *Graph[T]) Display() map[T][]T {
	return gp.adjList
}
