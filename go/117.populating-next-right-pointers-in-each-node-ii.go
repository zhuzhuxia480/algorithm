/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
package main
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0, 100)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLen := len(queue)
		preNode := queue[0]
		for i := 0; i < curLen; i++ {
			if i > 0 {
				preNode.Next = queue[i]
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
	
			}
			preNode = queue[i]
		}
		queue = queue[curLen:]
	}
	return root
}
// @lc code=end

