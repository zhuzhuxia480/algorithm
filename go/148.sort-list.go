/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	slow, quick := head, head.Next
	for quick != nil && quick.Next != nil{
		quick = quick.Next.Next
		slow = slow.Next
	}
	right := slow.Next
	slow.Next = nil
	leftHead := sortList(head)
	rightHead := sortList(right)

	dummy := &ListNode{}
	pTemp := dummy
	pleft, pright := leftHead, rightHead
	for pleft != nil && pright != nil {
		if pleft.Val < pright.Val {
			pTemp.Next = pleft
			pleft = pleft.Next
		} else {
			pTemp.Next = pright
			pright = pright.Next
		}
		pTemp = pTemp.Next
	}

	if pleft != nil {
		pTemp.Next = pleft
	}

	if pright != nil {
		pTemp.Next = pright
	}
	return dummy.Next
}
// @lc code=end

