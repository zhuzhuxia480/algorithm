/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
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
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	pPre := (*ListNode)(nil)
// 	pCur := head
// 	pNext := pCur.Next
	
// 	for pCur != nil {
// 		pCur.Next = pPre
// 		pPre = pCur
// 		pCur = pNext
// 		if pCur == nil {
// 			return pPre
// 		}
// 		pNext = pCur.Next
// 	}	
// 	return head
// }

var newHead *ListNode
func dfs(node *ListNode) {
	if node.Next == nil {
		newHead = node
		return
	}
	dfs(node.Next)
	node.Next.Next = node
}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dfs(head)
	head.Next = nil
	return newHead
}
// @lc code=end

