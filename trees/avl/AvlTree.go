package avl

type numericTypes interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type AvlNode[T numericTypes] struct {
	Left   *AvlNode[T]
	Right  *AvlNode[T]
	Data   T
	Height int
}

type AvlTree[T numericTypes] struct {
	Root *AvlNode[T]
}

// NewAvlTree() -> creates an AVL tree
func NewAvlTree[T numericTypes]() *AvlTree[T] {
	return &AvlTree[T]{}
}

// height() ->  0 if root
func height[T numericTypes](n *AvlNode[T]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// updateHeight()
func updateHeight[T numericTypes](node *AvlNode[T]) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

// balanceFactor() -> only 0 , -1 ,-2 ok
func balanceFactor[T numericTypes](n *AvlNode[T]) int {
	if n == nil {
		return 0
	}

	return height(n.Left) - height(n.Right)
}

// rotateRight() -> fixes a left heavy subtree, returns the new subtree root
func rotateRight[T numericTypes](y *AvlNode[T]) *AvlNode[T] {
	x := y.Left
	b := x.Right

	x.Right = y
	y.Left = b

	updateHeight(y)
	updateHeight(x)

	return x
}

// rotateLeft() -> fixes a right heavy subtree, returns the new subtree root
func rotateLeft[T numericTypes](x *AvlNode[T]) *AvlNode[T] {
	y := x.Right
	b := y.Left

	y.Left = x
	x.Right = b

	updateHeight(x)
	updateHeight(y)

	return y
}

// Insert() -> inserts data, keeping the tree balanced
func (tree *AvlTree[T]) Insert(data T) {
	tree.Root = insertHelper(tree.Root, data)
}

func insertHelper[T numericTypes](node *AvlNode[T], data T) *AvlNode[T] {
	//  normal BST insert
	if node == nil {
		return &AvlNode[T]{Data: data, Height: 1}
	}

	if data < node.Data {
		node.Left = insertHelper(node.Left, data)
	} else if data > node.Data {
		node.Right = insertHelper(node.Right, data)
	} else {
		return node // duplicate, ignore
	}

	updateHeight(node)

	bf := balanceFactor(node)

	if bf > 1 {
		if data > node.Left.Data { // Left-Right case
			node.Left = rotateLeft(node.Left)
		}
		return rotateRight(node) // Left-Left case
	}

	if bf < -1 {
		if data < node.Right.Data {
			node.Right = rotateRight(node.Right)
		}
		return rotateLeft(node)
	}

	return node
}

// Find() => binary search implementation returns a pointer to the node and true if val is found in tree, else returns zero struct and false
func (t AvlTree[T]) Find(val T) (*AvlNode[T], bool) {
	zero := &AvlNode[T]{}

	if t.Root == nil {
		return zero, false
	}

	for t.Root != nil {
		if val == t.Root.Data {
			return t.Root, true
		} else if val < t.Root.Data {
			t.Root = t.Root.Left
		} else if val > t.Root.Data {
			t.Root = t.Root.Right
		}
	}

	return zero, false
}
