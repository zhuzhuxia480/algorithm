/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	numLen := len(nums)
	if numLen == 1 {
		return &TreeNode{
			Val: nums[0],
			Left: nil,
			Right: nil,
		}
	}
	mid := numLen/2
	root := &TreeNode{
		Val: nums[mid],
		Right: nil,
		Left: nil,
	}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[0:mid])
	}
	if mid < numLen - 1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}
// @lc code=end

