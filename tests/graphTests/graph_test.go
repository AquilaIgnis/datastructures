package tests

import (
	"slices"
	"testing"

	"datastructures/graphs"
)

func TestNewGraphIsEmpty(t *testing.T) {
	graph := graphs.NewGraph[int]()

	if graph.Order() != 0 {
		t.Errorf("new graph Order = %d, want 0", graph.Order())
	}
	if graph.Size() != 0 {
		t.Errorf("new graph Size = %d, want 0", graph.Size())
	}
}

func TestAddVertexCreatesBothAndConnects(t *testing.T) {
	graph := graphs.NewGraph[string]()
	graph.AddVertex("a", "b")

	if graph.Order() != 2 {
		t.Errorf("Order = %d, want 2", graph.Order())
	}
	if graph.Size() != 1 {
		t.Errorf("Size = %d, want 1", graph.Size())
	}

	adjacency := graph.Display()
	if !slices.Contains(adjacency["a"], "b") {
		t.Errorf("expected b in a's adjacency, got %v", adjacency["a"])
	}
	if !slices.Contains(adjacency["b"], "a") {
		t.Errorf("expected a in b's adjacency (undirected), got %v", adjacency["b"])
	}
}

func TestAddVertexIsIdempotentForEdges(t *testing.T) {
	graph := graphs.NewGraph[string]()
	graph.AddVertex("a", "b")
	graph.AddVertex("a", "b")
	graph.AddVertex("b", "a")

	adjacency := graph.Display()
	if len(adjacency["a"]) != 1 {
		t.Errorf("expected no duplicate edges in a, got %v", adjacency["a"])
	}
	if len(adjacency["b"]) != 1 {
		t.Errorf("expected no duplicate edges in b, got %v", adjacency["b"])
	}
	if graph.Size() != 1 {
		t.Errorf("Size = %d, want 1", graph.Size())
	}
}

func TestAddVertexMultipleEdges(t *testing.T) {
	graph := graphs.NewGraph[int]()
	graph.AddVertex(1, 2)
	graph.AddVertex(1, 3)
	graph.AddVertex(1, 4)

	if graph.Order() != 4 {
		t.Errorf("Order = %d, want 4", graph.Order())
	}
	if graph.Size() != 3 {
		t.Errorf("Size = %d, want 3", graph.Size())
	}
	if len(graph.Display()[1]) != 3 {
		t.Errorf("expected vertex 1 to have degree 3, got %v", graph.Display()[1])
	}
}

func TestAddIsolatedVertex(t *testing.T) {
	graph := graphs.NewGraph[string]()

	if !graph.AddIsolatedVertex("solo") {
		t.Error("AddIsolatedVertex on new vertex = false, want true")
	}
	if graph.AddIsolatedVertex("solo") {
		t.Error("AddIsolatedVertex on existing vertex = true, want false")
	}

	if graph.Order() != 1 {
		t.Errorf("Order = %d, want 1", graph.Order())
	}
	if graph.Size() != 0 {
		t.Errorf("isolated vertex should add no edges, Size = %d", graph.Size())
	}
	if len(graph.Display()["solo"]) != 0 {
		t.Errorf("isolated vertex should have empty adjacency, got %v", graph.Display()["solo"])
	}
}

func TestConnectEdgesAllExisting(t *testing.T) {
	graph := graphs.NewGraph[int]()
	graph.AddIsolatedVertex(1)
	graph.AddIsolatedVertex(2)
	graph.AddIsolatedVertex(3)
	graph.AddIsolatedVertex(4)

	if !graph.ConnectEdges(1, 2, 3, 4) {
		t.Error("ConnectEdges with all existing vertices = false, want true")
	}
	if graph.Size() != 3 {
		t.Errorf("Size = %d, want 3", graph.Size())
	}
	if len(graph.Display()[1]) != 3 {
		t.Errorf("expected vertex 1 to connect to 3 vertices, got %v", graph.Display()[1])
	}
	// undirected: each target should point back at 1
	for _, target := range []int{2, 3, 4} {
		if !slices.Contains(graph.Display()[target], 1) {
			t.Errorf("expected vertex %d to point back at 1, got %v", target, graph.Display()[target])
		}
	}
}

func TestConnectEdgesRejectsMissingVertex(t *testing.T) {
	graph := graphs.NewGraph[int]()
	graph.AddIsolatedVertex(1)
	graph.AddIsolatedVertex(2)
	// 99 does not exist

	if graph.ConnectEdges(1, 2, 99) {
		t.Error("ConnectEdges with a missing target = true, want false")
	}
	// nothing should have been connected (atomic)
	if graph.Size() != 0 {
		t.Errorf("no edges should be added on failure, Size = %d", graph.Size())
	}
	if len(graph.Display()[1]) != 0 {
		t.Errorf("vertex 1 should remain isolated on failure, got %v", graph.Display()[1])
	}
}

func TestSelfLoopIsRejected(t *testing.T) {
	graph := graphs.NewGraph[int]()
	graph.AddVertex(1, 1)

	if graph.Size() != 0 {
		t.Errorf("self-loop should add no edge, Size = %d", graph.Size())
	}
	if slices.Contains(graph.Display()[1], 1) {
		t.Errorf("vertex 1 should not contain itself, got %v", graph.Display()[1])
	}

	// a real edge on the same vertex must still work
	graph.AddVertex(1, 2)
	if graph.Size() != 1 {
		t.Errorf("Size = %d, want 1 after a valid edge", graph.Size())
	}
}

func TestOrderAndSize(t *testing.T) {
	graph := graphs.NewGraph[string]()
	graph.AddVertex("a", "b")
	graph.AddVertex("b", "c")
	graph.AddVertex("c", "a")

	if graph.Order() != 3 {
		t.Errorf("Order = %d, want 3 (triangle)", graph.Order())
	}
	if graph.Size() != 3 {
		t.Errorf("Size = %d, want 3 (triangle)", graph.Size())
	}
}
