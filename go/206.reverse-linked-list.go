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
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pPre := (*ListNode)(nil)
	pCur := head
	pNext := pCur.Next
	
	for pCur != nil {
		pCur.Next = pPre
		pPre = pCur
		pCur = pNext
		if pCur == nil {
			return pPre
		}
		pNext = pCur.Next
	}	
	return head
}
// @lc code=end

