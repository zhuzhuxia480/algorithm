/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main
func countDepth(root *TreeNode) int {
	depth := 0
	for root != nil {
		depth++
		root = root.Left
	}
	return depth
}
func countNodes(root *TreeNode) int {
    if root == nil {
		return 0
	}
	leftDepth := countDepth(root.Left)
	rightDepth := countDepth(root.Right)
	if leftDepth == rightDepth {
		return (1<<leftDepth) + countNodes(root.Right)
	} else {
		return (1<<rightDepth) + countNodes(root.Left)
	}
}
// @lc code=end

