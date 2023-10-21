package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	leftDepth := getMaxDepth(root.Left, depth+1)
	rightDepth := getMaxDepth(root.Right, depth+1)
	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}

func maxDepth(root *TreeNode) int {
	return getMaxDepth(root, 0)
}
