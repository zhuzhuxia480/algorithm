/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
var stop = false
var flag = true
func dsfBST(root *TreeNode) (int, int) {
	if stop {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	lmin, lmax, rmin, rmax := root.Val, root.Val, root.Val, root.Val
	if root.Left != nil {
		lmin, lmax = dsfBST(root.Left)
		if lmax >= root.Val {
			stop = true
			flag = false
		}
	}
	if root.Right != nil {
		rmin, rmax = dsfBST(root.Right)
		if rmin <= root.Val {
			stop = true
			flag = false
		}
	}
	return lmin, rmax
}
func isValidBST(root *TreeNode) bool {
    flag = true
	stop = false
	dsfBST(root)
	return flag
}
// @lc code=end

