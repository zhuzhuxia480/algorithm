/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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

func dfsAncestor(root *TreeNode, p, q *TreeNode, res **TreeNode) bool {
	if *res != nil {
		return false
	}

	if root == nil {
		return false
	}

	lFlag := dfsAncestor(root.Left, p, q, res)
	rFlag := dfsAncestor(root.Right, p, q, res)
	if lFlag && rFlag && *res == nil {
		*res = root
		return false
	}

	if root == p || root == q {
		if lFlag || rFlag && *res == nil {
			*res = root
			return false
		}
		return true
	}
	return lFlag || rFlag
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res := (*TreeNode)(nil)
	dfsAncestor(root, p, q, &res)
	return res
}
// @lc code=end

