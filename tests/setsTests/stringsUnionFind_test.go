package tests

import (
	"testing"

	"datastructures/sets"
)

// --- NewStringsUnion ---

func TestNewStringsUnionStartsFullyDisjointed(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c", "d"})
	if got := u.Disjointed(); got != 4 {
		t.Errorf("expected 4 disjoint sets initially, got %d", got)
	}
}

func TestNewStringsUnionDeduplicates(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "a", "c", "b", "a"})
	// only a, b, c are unique
	if got := u.Disjointed(); got != 3 {
		t.Errorf("expected 3 disjoint sets after dedup, got %d", got)
	}
}

func TestNewStringsUnionEachNameIsOwnRep(t *testing.T) {
	names := []string{"alice", "bob", "carol"}
	u := sets.NewStringsUnion(names)
	for _, name := range names {
		rep, ok := u.Rep(name)
		if !ok {
			t.Errorf("expected %q to be known", name)
		}
		if rep != name {
			t.Errorf("expected %q to be its own representative, got %q", name, rep)
		}
	}
}

func TestNewStringsUnionEmptyInput(t *testing.T) {
	u := sets.NewStringsUnion(nil)
	if got := u.Disjointed(); got != 0 {
		t.Errorf("expected 0 disjoint sets for empty input, got %d", got)
	}
}

// --- Union ---

func TestNamedUnionReturnsTrueOnMerge(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c"})
	if !u.Union("a", "b") {
		t.Error("expected Union(a,b) to return true on a fresh merge")
	}
}

func TestNamedUnionReturnsFalseWhenAlreadyJoined(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c"})
	u.Union("a", "b")
	if u.Union("a", "b") {
		t.Error("expected Union(a,b) to return false when already in the same set")
	}
}

func TestNamedUnionReturnsFalseTransitively(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c"})
	u.Union("a", "b")
	u.Union("b", "c")
	// a and c are now connected through b
	if u.Union("a", "c") {
		t.Error("expected Union(a,c) to return false (connected via b)")
	}
}

func TestNamedUnionUnknownNameReturnsFalse(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b"})
	if u.Union("a", "ghost") {
		t.Error("expected Union with an unknown name to return false")
	}
	// the unknown name must not have silently merged anything
	if got := u.Disjointed(); got != 2 {
		t.Errorf("expected the count to stay at 2 after a failed union, got %d", got)
	}
}

func TestNamedUnionDecrementsDisjointedCount(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c", "d"})
	u.Union("a", "b") // 4 -> 3
	u.Union("c", "d") // 3 -> 2
	if got := u.Disjointed(); got != 2 {
		t.Errorf("expected 2 disjoint sets, got %d", got)
	}
}

func TestNamedUnionIsSymmetric(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b"})
	u.Union("b", "a")
	if !u.IsUnion("a", "b") {
		t.Error("expected Union(b,a) to connect a and b regardless of argument order")
	}
}

// --- IsUnion ---

func TestNamedIsUnionTrueAfterMerge(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c"})
	u.Union("a", "b")
	if !u.IsUnion("a", "b") {
		t.Error("expected a and b to be in the same set after Union")
	}
}

func TestNamedIsUnionTransitive(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c", "d"})
	u.Union("a", "b")
	u.Union("b", "c")
	if !u.IsUnion("a", "c") {
		t.Error("expected a and c to be connected through b")
	}
	if u.IsUnion("a", "d") {
		t.Error("expected d to remain disjoint from the {a,b,c} set")
	}
}

func TestNamedIsUnionUnknownNameReturnsFalse(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b"})
	if u.IsUnion("a", "ghost") {
		t.Error("expected IsUnion with an unknown name to return false")
	}
}

// --- Rep ---

func TestRepIsSharedAfterUnion(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c"})
	u.Union("a", "b")
	u.Union("b", "c")

	repA, _ := u.Rep("a")
	repB, _ := u.Rep("b")
	repC, _ := u.Rep("c")
	if repA != repB || repB != repC {
		t.Errorf("expected a, b, c to share one representative, got %q %q %q", repA, repB, repC)
	}
}

func TestRepUnknownNameReturnsFalse(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b"})
	if rep, ok := u.Rep("ghost"); ok || rep != "" {
		t.Errorf("expected (\"\", false) for an unknown name, got (%q, %v)", rep, ok)
	}
}

func TestRepDistinctForSeparateSets(t *testing.T) {
	u := sets.NewStringsUnion([]string{"a", "b", "c", "d"})
	u.Union("a", "b")
	u.Union("c", "d")

	repA, _ := u.Rep("a")
	repC, _ := u.Rep("c")
	if repA == repC {
		t.Errorf("expected separate sets to have different representatives, both were %q", repA)
	}
}

// --- Scenario: grouping people into friend circles ---

func TestFriendCircles(t *testing.T) {
	friendships := [][2]string{
		{"alice", "bob"}, {"bob", "carol"}, // circle A: alice-bob-carol
		{"dave", "erin"},                    // circle B: dave-erin
		{"frank", "grace"}, {"grace", "frank"}, // circle C: frank-grace (redundant)
	}

	circles := sets.NewStringsUnion([]string{
		"alice", "bob", "carol", "dave", "erin", "frank", "grace",
	})
	for _, f := range friendships {
		circles.Union(f[0], f[1])
	}

	if got := circles.Disjointed(); got != 3 {
		t.Errorf("expected 3 friend circles, got %d", got)
	}
	if !circles.IsUnion("alice", "carol") {
		t.Error("expected alice and carol to be in the same circle")
	}
	if circles.IsUnion("alice", "dave") {
		t.Error("expected alice (circle A) and dave (circle B) to be in different circles")
	}
	if !circles.IsUnion("frank", "grace") {
		t.Error("expected frank and grace to be in the same circle despite the redundant link")
	}
}
