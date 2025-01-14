/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
var res = -1
func dfsKth(root *TreeNode, k int, index *int) {
	if res != -1 {
		return 
	}
	if root == nil {
		return 
	}
	dfsKth(root.Left, k, index)
	(*index)++
	if *index == k {
		res = root.Val
		return 
	}
	dfsKth(root.Right, k, index)
	
}
func kthSmallest(root *TreeNode, k int) int {
	res = -1
	index := 0
    dfsKth(root, k, &index)
	return res

}
// @lc code=end

