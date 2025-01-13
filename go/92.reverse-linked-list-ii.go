/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
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
func dfsNode(pStart *ListNode, pEnd *ListNode) {
	if pStart == pEnd {
		return
	}
	dfsNode(pStart.Next, pEnd)
	pStart.Next.Next = pStart
}
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right || head == nil || head.Next == nil {
		return head
	}
	pPre := (*ListNode)(nil)
	pLast := pPre
	pStart := head
	pEnd := pStart
	ptmp := head

	for i := 1; i <= right; i++ {
		if i == left-1 {
			pPre = ptmp
			pStart = ptmp.Next
		}
		if i == right {
			pEnd = ptmp
			pLast = ptmp.Next
		}
		ptmp = ptmp.Next
	}
	dfsNode(pStart, pEnd)
	if pPre != nil {
		pPre.Next = pEnd
	} else {
		head = pEnd
	}
	pStart.Next = pLast
	return head
}
// @lc code=end

