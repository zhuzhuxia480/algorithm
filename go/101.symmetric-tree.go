/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isMirrorTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			return isMirrorTree(p.Left, q.Right) && isMirrorTree(p.Right, q.Left)
		} else {
			return false
		}
	}
	return false
} 
func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return true
	}
	return isMirrorTree(root.Left, root.Right)
}
// @lc code=end

