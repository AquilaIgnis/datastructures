package trees

// Binary search tree

import (
	"fmt"
)

type NumericTypes interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Node[T NumericTypes] struct {
	left  *Node[T]
	right *Node[T]
	Data  T
}

// BSTree() -> Binary search tree, &BSTree[type] , must be in NumericTypes ;
type BSTree[T NumericTypes] struct {
	root *Node[T]
}

// Insert() -> Inserts Node
func (tree *BSTree[T]) Insert(n T) {
	tree.root = insertHelper(tree.root, n)
}

func insertHelper[T NumericTypes](node *Node[T], n T) *Node[T] {
	if node == nil {
		return &Node[T]{Data: n}
	}

	switch {
	case n < node.Data:
		node.left = insertHelper(node.left, n)

	case n > node.Data:
		node.right = insertHelper(node.right, n)
	}

	return node
}

// Display() -> prints nodes :: recursive implementation
func (t *BSTree[T]) Display() {
	displayHelper(t.root)
}

func displayHelper[T NumericTypes](node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Println("node:", node.Data)
	if node.left != nil {
		fmt.Println("  left child:", node.left.Data)
	}
	if node.right != nil {
		fmt.Println("  right child:", node.right.Data)
	}

	displayHelper(node.left)
	displayHelper(node.right)
}

// Find() => binary search implementation returns a pointer to the node and true if val is found in tree, else returns zero struct and false
func (t BSTree[T]) Find(val T) (*Node[T], bool) {
	zero := &Node[T]{}

	if t.root == nil {
		return zero, false
	}

	for t.root != nil {
		if val == t.root.Data {
			return t.root, true
		} else if val < t.root.Data {
			t.root = t.root.left
		} else if val > t.root.Data {
			t.root = t.root.right
		}
	}

	fmt.Println("Data not Found")
	return zero, false
}
