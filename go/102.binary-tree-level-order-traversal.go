/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLen := len(queue)
		tmpRes := []int{}
		for i := 0; i < curLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmpRes = append(tmpRes, queue[i].Val)
		}
		queue = queue[curLen:]
		res = append(res, tmpRes)
	}
	return res
}
// @lc code=end

