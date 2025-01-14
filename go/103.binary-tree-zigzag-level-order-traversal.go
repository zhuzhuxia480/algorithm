/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := [][]int{}
	queue := []*TreeNode{root}
	if root == nil {
		return res
	}
	flag := 0
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
			if flag%2 == 0 {
				tmpRes = append(tmpRes, queue[i].Val)
			} else {
				tmpRes = append(tmpRes, queue[curLen-i-1].Val)
			}
		}
		queue = queue[curLen:]
		res = append(res, tmpRes)
		flag++
	}
	return res
}
// @lc code=end

