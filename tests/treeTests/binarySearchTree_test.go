package tests

import (
	"bufio"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"datastructures/trees"
)

// The BSTree exposes only Insert and Display publicly; root/left/right are
// unexported, so these tests drive the tree through Insert and observe it
// through the pre-order output that Display writes to stdout.

// captureDisplay runs t.Display() while capturing stdout and returns the
// pre-order sequence of node values (the "node:" lines), ignoring the
// "left child"/"right child" annotation lines.
func captureDisplay(tree *trees.BSTree[int]) []int {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tree.Display()

	w.Close()
	os.Stdout = old

	var pre []int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		rest, ok := strings.CutPrefix(line, "node:")
		if !ok {
			continue
		}
		v, err := strconv.Atoi(strings.TrimSpace(rest))
		if err != nil {
			continue
		}
		pre = append(pre, v)
	}
	_, _ = io.Copy(io.Discard, r)
	return pre
}

// validBSTPreorder reports whether pre is a valid pre-order traversal of a BST
// with distinct keys (classic monotonic-stack check).
func validBSTPreorder(pre []int) bool {
	stack := []int{}
	lower := math.MinInt
	for _, v := range pre {
		if v <= lower {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] < v {
			lower = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	return true
}

func sortedCopy(in []int) []int {
	out := append([]int(nil), in...)
	sort.Ints(out)
	return out
}

func uniqueSorted(in []int) []int {
	set := map[int]struct{}{}
	for _, v := range in {
		set[v] = struct{}{}
	}
	out := make([]int, 0, len(set))
	for v := range set {
		out = append(out, v)
	}
	sort.Ints(out)
	return out
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInsertMaintainsBSTInvariant(t *testing.T) {
	inputs := [][]int{
		{5, 3, 8, 1, 4, 7, 9},
		{1, 2, 3, 4, 5},          // degenerate right-leaning
		{5, 4, 3, 2, 1},          // degenerate left-leaning
		{42},                     // single node
		{10, 5, 15, 3, 7, 13, 17, 1, 4, 6, 8},
	}

	for _, in := range inputs {
		tree := &trees.BSTree[int]{}
		for _, v := range in {
			tree.Insert(v)
		}

		pre := captureDisplay(tree)

		if !validBSTPreorder(pre) {
			t.Errorf("inserting %v produced invalid BST pre-order %v", in, pre)
		}
		// Every inserted value should be present exactly once.
		if got, want := sortedCopy(pre), uniqueSorted(in); !equalInts(got, want) {
			t.Errorf("inserting %v: stored values %v, want %v", in, got, want)
		}
	}
}

func TestInsertIgnoresDuplicates(t *testing.T) {
	tree := &trees.BSTree[int]{}
	for _, v := range []int{5, 3, 5, 8, 3, 5, 8, 8} {
		tree.Insert(v)
	}

	pre := captureDisplay(tree)
	if got, want := sortedCopy(pre), []int{3, 5, 8}; !equalInts(got, want) {
		t.Errorf("duplicates not ignored: got %v, want %v", got, want)
	}
}

func TestInsertRootFirst(t *testing.T) {
	tree := &trees.BSTree[int]{}
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)

	pre := captureDisplay(tree)
	// Pre-order visits the root first.
	if len(pre) == 0 || pre[0] != 50 {
		t.Errorf("expected root 50 first in pre-order, got %v", pre)
	}
}

func TestDisplayEmptyTree(t *testing.T) {
	tree := &trees.BSTree[int]{}
	pre := captureDisplay(tree)
	if len(pre) != 0 {
		t.Errorf("empty tree should produce no nodes, got %v", pre)
	}
}

func TestFindReturnsNodeForPresentValues(t *testing.T) {
	values := []int{50, 25, 75, 10, 30, 60, 90}
	tree := &trees.BSTree[int]{}
	for _, v := range values {
		tree.Insert(v)
	}

	for _, v := range values {
		node, ok := tree.Find(v)
		if !ok {
			t.Errorf("Find(%d): ok=false, want true", v)
			continue
		}
		if node == nil {
			t.Errorf("Find(%d): nil node with ok=true", v)
			continue
		}
		if node.Data != v {
			t.Errorf("Find(%d): node.Data=%d, want %d", v, node.Data, v)
		}
	}
}

func TestFindReturnsFalseForMissingValues(t *testing.T) {
	tree := &trees.BSTree[int]{}
	for _, v := range []int{50, 25, 75} {
		tree.Insert(v)
	}

	for _, v := range []int{0, 26, 100, -5} {
		if _, ok := tree.Find(v); ok {
			t.Errorf("Find(%d): ok=true, want false", v)
		}
	}
}

func TestFindOnEmptyTree(t *testing.T) {
	tree := &trees.BSTree[int]{}
	if _, ok := tree.Find(42); ok {
		t.Errorf("Find on empty tree: ok=true, want false")
	}
}

// TestFindDoesNotMutateTree guards against the traversal walking the tree with
// the root pointer itself: a lookup must leave the tree fully intact.
func TestFindDoesNotMutateTree(t *testing.T) {
	values := []int{50, 25, 75, 10, 30, 60, 90}
	tree := &trees.BSTree[int]{}
	for _, v := range values {
		tree.Insert(v)
	}

	before := captureDisplay(tree)

	// Look up every value plus some misses.
	for _, v := range append(append([]int(nil), values...), 0, 100, 55) {
		tree.Find(v)
	}

	after := captureDisplay(tree)
	if !equalInts(before, after) {
		t.Errorf("Find mutated the tree: pre-order before=%v, after=%v", before, after)
	}
}

func TestInsertOrderIndependence(t *testing.T) {
	// Different insertion orders of the same set must store the same values.
	a := &trees.BSTree[int]{}
	for _, v := range []int{4, 2, 6, 1, 3, 5, 7} {
		a.Insert(v)
	}
	b := &trees.BSTree[int]{}
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
		b.Insert(v)
	}

	if got, want := sortedCopy(captureDisplay(a)), sortedCopy(captureDisplay(b)); !equalInts(got, want) {
		t.Errorf("same set, different order stored differently: %v vs %v", got, want)
	}
}
