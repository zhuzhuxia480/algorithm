/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
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
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil  {
		return head
	}
	length := 1
	pTail := head
	for pTail.Next != nil {
		length++
		pTail = pTail.Next
	}
	if k%length == 0 {
		return head
	}
	pHead := head
	pPre := head
	index := length - k%length
	pCur := head
	for i := 0; i < index; i++ {
		if i == index-1 {
			pPre = pCur
		}
		pCur = pCur.Next
	}
	pHead = pCur
	pPre.Next = nil
	pTail.Next = head
	return pHead
}
// @lc code=end

