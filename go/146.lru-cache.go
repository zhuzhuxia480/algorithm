/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
package main

type LinkedListNode struct {
	Val  int
	Key  int
	Next *LinkedListNode
	Pre  *LinkedListNode
}
type LRUCache struct {
	capacity int
	length   int
	head     *LinkedListNode
	nodeMap  map[int]*LinkedListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		length:   0,
		head: &LinkedListNode{
			Val:  -1,
			Pre:  nil,
			Next: nil,
		},
		nodeMap: make(map[int]*LinkedListNode),
	}
	lru.head.Next = lru.head
	lru.head.Pre = lru.head.Pre
	return lru
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.nodeMap[key]
	if !ok {
		return -1
	}

	this.moveToHead(v)
	return v.Val
}

func (this *LRUCache) moveToHead(node *LinkedListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	node.Next = this.head.Next
	this.head.Next.Pre = node
	this.head.Next = node
	node.Pre = this.head
}

func (this *LRUCache) deleteOne() {
	node := this.head.Pre
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	delete(this.nodeMap, node.Key)
	this.length--
}

func (this *LRUCache) Put(key int, value int) {

	v, ok := this.nodeMap[key]
	if ok {
		v.Val = value
		this.moveToHead(v)
	} else {
		if this.length == this.capacity {
			this.deleteOne()
		}
		newNode := &LinkedListNode{
			Val:  value,
			Key:  key,
			Pre:  this.head,
			Next: this.head.Next,
		}
		this.head.Next.Pre = newNode
		this.head.Next = newNode
		this.nodeMap[key] = newNode
		this.length++
		if this.length == 1 {
			this.head.Pre = newNode
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
