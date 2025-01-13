/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
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
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
    pLessHead := &ListNode{0, nil}
	pGreaterHead :=&ListNode{0, nil}
	pCur := head
	pLessCur := pLessHead
	pGreaterCur := pGreaterHead
	for pCur != nil {
		if pCur.Val < x {
			pLessCur.Next = pCur
			pLessCur = pCur
		} else {
			pGreaterCur.Next = pCur
			pGreaterCur = pCur
		}
		pCur = pCur.Next
	}
	pLessCur.Next = pGreaterHead.Next
	pGreaterCur.Next = nil
	return pLessHead.Next
}
// @lc code=end

