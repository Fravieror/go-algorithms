package main

import "fmt"

func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("Inorder Traversal (Recursive):")
	InorderTraversal(root)
	fmt.Println()

	fmt.Println("Preorder Traversal (Recursive):")
	PreorderTraversal(root)
	fmt.Println()

	fmt.Println("Postorder Traversal (Recursive):")
	PostorderTraversal(root)
	fmt.Println()

	fmt.Println("Inorder Traversal (Non-Recursive):")
	InorderTraversalNonRecursive(root)
	fmt.Println()

	fmt.Println("Preorder Traversal (Non-Recursive):")
	PreorderTraversalNonRecursive(root)
	fmt.Println()

	fmt.Println("Postorder Traversal (Non-Recursive):")
	PostorderTraversalNonRecursive(root)
	fmt.Println()

	fmt.Println("Level Order Traversal (Breadth-First):")
	LevelOrderTraversal(root)
	fmt.Println()
}

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert inserts a new value into the binary tree.
func (t *TreeNode) Insert(value int) {
	if value < t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: value}
		} else {
			t.Left.Insert(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: value}
		} else {
			t.Right.Insert(value)
		}
	}
}
