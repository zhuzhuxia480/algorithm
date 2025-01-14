/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func dfsFlatten(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	lLast := dfsFlatten(root.Left)
	rLast := dfsFlatten(root.Right)

	if lLast == nil {
		root.Left = nil
		return rLast
	}
	if rLast == nil {
		root.Right = root.Left
		root.Left = nil
		return lLast
	}
	if rLast != nil && lLast != nil {
		lLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		return rLast
	}
	return rLast
}
func flatten(root *TreeNode)  {
    if root == nil {
		return
	}
	dfsFlatten(root)
}
// @lc code=end

