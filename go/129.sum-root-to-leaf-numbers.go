/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func dfsSum(root *TreeNode, currentVal int, totalSum *int) {
	currentVal = currentVal*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*totalSum += currentVal
	}
	if root.Left != nil {
		dfsSum(root.Left, currentVal, totalSum)
	}
	if root.Right != nil {
		dfsSum(root.Right, currentVal, totalSum)
	}
}
func sumNumbers(root *TreeNode) int {
    totalSum := 0
	dfsSum(root, 0, &totalSum)
	return totalSum
}
// @lc code=end

