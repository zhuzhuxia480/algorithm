/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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

func checkSum(root *TreeNode, targetSum int, tmpSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if tmpSum + root.Val == targetSum {
			return true
		}
		return false
	}
	return checkSum(root.Left, targetSum, tmpSum+root.Val) || checkSum(root.Right, targetSum, tmpSum+root.Val)
}
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return checkSum(root, targetSum, 0)
}
// @lc code=end

