package tests

import (
	"testing"

	"datastructures/sets"
)

// --- NewUnionFind ---

func TestNewUnionFindStartsFullyDisjointed(t *testing.T) {
	uf := sets.NewUnionFind(8)
	if got := uf.Disjointed(); got != 8 {
		t.Errorf("expected 8 disjoint sets initially, got %d", got)
	}
}

func TestNewUnionFindEachNodeIsOwnRoot(t *testing.T) {
	uf := sets.NewUnionFind(5)
	for i := 1; i <= 5; i++ {
		if root := uf.Find(i); root != i {
			t.Errorf("expected node %d to be its own root, got %d", i, root)
		}
	}
}

func TestNewUnionFindNothingIsUnionedYet(t *testing.T) {
	uf := sets.NewUnionFind(5)
	for a := 1; a <= 5; a++ {
		for b := 1; b <= 5; b++ {
			if a == b {
				continue
			}
			if uf.IsUnion(a, b) {
				t.Errorf("expected %d and %d to be disjoint initially", a, b)
			}
		}
	}
}

// --- Union ---

func TestUnionReturnsTrueOnMerge(t *testing.T) {
	uf := sets.NewUnionFind(3)
	if !uf.Union(1, 2) {
		t.Error("expected Union(1,2) to return true on a fresh merge")
	}
}

func TestUnionReturnsFalseWhenAlreadyJoined(t *testing.T) {
	uf := sets.NewUnionFind(3)
	uf.Union(1, 2)
	if uf.Union(1, 2) {
		t.Error("expected Union(1,2) to return false when already in the same set")
	}
}

func TestUnionReturnsFalseTransitively(t *testing.T) {
	uf := sets.NewUnionFind(3)
	uf.Union(1, 2)
	uf.Union(2, 3)
	// 1 and 3 are now connected through 2
	if uf.Union(1, 3) {
		t.Error("expected Union(1,3) to return false (connected via 2)")
	}
}

func TestUnionDecrementsDisjointedCount(t *testing.T) {
	uf := sets.NewUnionFind(4)
	uf.Union(1, 2) // 4 -> 3
	uf.Union(3, 4) // 3 -> 2
	if got := uf.Disjointed(); got != 2 {
		t.Errorf("expected 2 disjoint sets, got %d", got)
	}
}

func TestUnionRedundantDoesNotDecrement(t *testing.T) {
	uf := sets.NewUnionFind(4)
	uf.Union(1, 2)
	uf.Union(1, 2) // redundant, must not change the count
	if got := uf.Disjointed(); got != 3 {
		t.Errorf("expected 3 disjoint sets after a redundant union, got %d", got)
	}
}

func TestUnionIsSymmetric(t *testing.T) {
	uf := sets.NewUnionFind(2)
	uf.Union(2, 1)
	if !uf.IsUnion(1, 2) {
		t.Error("expected Union(2,1) to connect 1 and 2 regardless of argument order")
	}
}

// --- IsUnion / Find ---

func TestIsUnionTrueAfterMerge(t *testing.T) {
	uf := sets.NewUnionFind(3)
	uf.Union(1, 2)
	if !uf.IsUnion(1, 2) {
		t.Error("expected 1 and 2 to be in the same set after Union")
	}
}

func TestIsUnionTransitive(t *testing.T) {
	uf := sets.NewUnionFind(4)
	uf.Union(1, 2)
	uf.Union(2, 3)
	if !uf.IsUnion(1, 3) {
		t.Error("expected 1 and 3 to be connected through 2")
	}
	if uf.IsUnion(1, 4) {
		t.Error("expected 4 to remain disjoint from the {1,2,3} set")
	}
}

func TestFindSharedRootAfterChain(t *testing.T) {
	uf := sets.NewUnionFind(5)
	// build a chain 1-2-3-4-5
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 4)
	uf.Union(4, 5)

	root := uf.Find(1)
	for i := 2; i <= 5; i++ {
		if uf.Find(i) != root {
			t.Errorf("expected node %d to share root %d, got %d", i, root, uf.Find(i))
		}
	}
}

func TestFindIsIdempotent(t *testing.T) {
	uf := sets.NewUnionFind(4)
	uf.Union(1, 2)
	uf.Union(2, 3)
	first := uf.Find(1)
	// path compression runs on the first call; the answer must not drift
	if second := uf.Find(1); first != second {
		t.Errorf("expected Find to be stable, got %d then %d", first, second)
	}
}

// --- Scenario: the island/network example from main.go ---

func TestNetworkIslands(t *testing.T) {
	cables := [][2]int{
		{0, 1}, {1, 2}, // island A: 0-1-2
		{3, 4},                 // island B: 3-4
		{5, 6}, {6, 7}, {5, 7}, // island C: 5-6-7 (redundant cable)
	}

	network := sets.NewUnionFind(8)
	for _, c := range cables {
		network.Union(c[0], c[1])
	}

	if got := network.Disjointed(); got != 3 {
		t.Errorf("expected 3 islands, got %d", got)
	}
	if !network.IsUnion(0, 2) {
		t.Error("expected 0 and 2 to be on the same island")
	}
	if network.IsUnion(0, 4) {
		t.Error("expected 0 (island A) and 4 (island B) to be on different islands")
	}
	if !network.IsUnion(5, 7) {
		t.Error("expected 5 and 7 to be on the same island despite the redundant cable")
	}
}
