/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    queue := []*TreeNode{}
	queue = append(queue, root)
	res := []int{}
	for len(queue) > 0 {
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[curLen-1].Val)
		queue = queue[curLen:]
	}	
	return res
}
// @lc code=end

