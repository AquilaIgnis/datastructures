package algo

import (
	"fmt"

	"datastructures/linear"
	"datastructures/trees"
)

// BstTree() => Breadth first search implementation for Binary trees O(n) ,
// usage var mytree trees.BSTree[[int]] ,
// BstTree(mytree);
func BstTree[T trees.NumericTypes](node trees.BSTree[T]) {
	if node.Root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := linear.Queue[*trees.Node[T]]{}
	depth := 0

	queue.Push(node.Root)

	for queue.Size() > 0 {

		// all nodes currently queued at a level = breadth
		breadth := queue.Size()

		// will run same amount as current queue size
		for range breadth {

			node, err := queue.Pop()
			if err != nil {
				break
			}

			data := node.Data
			fmt.Printf("depth :%v value: %v\n", depth, data)

			if node.Left != nil {
				queue.Push(node.Left)
			}

			if node.Right != nil {
				queue.Push(node.Right)
			}
		}

		depth++
	}
}
