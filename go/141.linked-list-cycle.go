/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package main
func hasCycle(head *ListNode) bool {
    nodeMap := make(map[*ListNode]bool)
	
	p := head
	for p != nil {
		_, ok := nodeMap[p] 
		if !ok {
			nodeMap[p] = true
		} else {
			return true
		}
		p = p.Next
	}
	return false
}
// @lc code=end

