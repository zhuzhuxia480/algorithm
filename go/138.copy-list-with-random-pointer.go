/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
package main




func copyNode(node *Node) *Node {
	return &Node{
		Val: node.Val,
		Next: nil,
		Random: nil,
	}
}
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	nodeMap := make(map[*Node]*Node)
    newHead := copyNode(head)
	nodeMap[head] = newHead
	pNewNodeLast := newHead
	pNode := head.Next
	for pNode != nil {
		newNode := copyNode(pNode)
		pNewNodeLast.Next = newNode
		pNewNodeLast = newNode
		nodeMap[pNode] = newNode
		pNode = pNode.Next
	}

	pNode = head
	for pNode != nil {
		if pNode.Random != nil {
			nodeMap[pNode].Random = nodeMap[pNode.Random]
		}
		pNode = pNode.Next
	}
	return newHead
}
// @lc code=end

