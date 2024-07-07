package main

import "fmt"

// InorderTraversal performs an inorder traversal of the binary tree.
func InorderTraversal(node *TreeNode) {
	if node != nil {
		InorderTraversal(node.Left)
		fmt.Print(node.Value, " ")
		InorderTraversal(node.Right)
	}
}

// PreorderTraversal performs a preorder traversal of the binary tree.
func PreorderTraversal(node *TreeNode) {
	if node != nil {
		fmt.Print(node.Value, " ")
		PreorderTraversal(node.Left)
		PreorderTraversal(node.Right)
	}
}

// PostorderTraversal performs a postorder traversal of the binary tree.
func PostorderTraversal(node *TreeNode) {
	if node != nil {
		PostorderTraversal(node.Left)
		PostorderTraversal(node.Right)
		fmt.Print(node.Value, " ")
	}
}
