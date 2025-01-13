/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
    pHead := &ListNode{-999999, head}
	pCur := head
	pPre := pHead
	for pCur != nil {
		flag := false
		for pCur.Next!=nil && pCur.Val == pCur.Next.Val {
			pCur = pCur.Next
			flag = true
		}
		if flag {
			pPre.Next = pCur.Next
			pCur = pCur.Next
		} else {
			pPre = pCur
			pCur = pCur.Next
		}

	}
	return pHead.Next
}
// @lc code=end

