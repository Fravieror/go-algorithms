package main

import (
	"container/list"
	"fmt"
)

// LevelOrderTraversal performs a level order traversal of the binary tree.
func LevelOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(*TreeNode)
		fmt.Print(node.Value, " ")

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}
