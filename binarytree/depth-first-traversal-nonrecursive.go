package main

import (
	"container/list"
	"fmt"
)

// InorderTraversalNonRecursive performs a non-recursive inorder traversal of the binary tree.
func InorderTraversalNonRecursive(root *TreeNode) {
	stack := list.New()
	node := root

	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

		element := stack.Back()
		stack.Remove(element)
		node = element.Value.(*TreeNode)
		fmt.Print(node.Value, " ")
		node = node.Right
	}
}

// PreorderTraversalNonRecursive performs a non-recursive preorder traversal of the binary tree.
func PreorderTraversalNonRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		node := element.Value.(*TreeNode)
		fmt.Print(node.Value, " ")

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
}

// PostorderTraversalNonRecursive performs a non-recursive postorder traversal of the binary tree.
func PostorderTraversalNonRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	var prev *TreeNode
	stack.PushBack(root)

	for stack.Len() > 0 {
		current := stack.Back().Value.(*TreeNode)

		if prev == nil || prev.Left == current || prev.Right == current {
			if current.Left != nil {
				stack.PushBack(current.Left)
			} else if current.Right != nil {
				stack.PushBack(current.Right)
			}
		} else if current.Left == prev {
			if current.Right != nil {
				stack.PushBack(current.Right)
			}
		} else {
			fmt.Print(current.Value, " ")
			stack.Remove(stack.Back())
		}

		prev = current
	}
}
