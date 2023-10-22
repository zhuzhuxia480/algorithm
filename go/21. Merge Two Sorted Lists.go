package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	p := &head
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		} else {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		}
	}
	for list1 != nil {
		p.Next = list1
		list1 = list1.Next
		p = p.Next
	}
	for list2 != nil {
		p.Next = list2
		list2 = list2.Next
		p = p.Next
	}
	return head.Next
}
func main() {

}
