/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"
type BSTIterator struct {
    data []int
}

func inorderNode(data *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorderNode(data, root.Left)
	*data = append(*data, root.Val)
	inorderNode(data, root.Right)
}

func Constructor(root *TreeNode) BSTIterator {
	array := []int{}
	inorderNode(&array, root)
    return BSTIterator{
		data: array,
	}
}


func (this *BSTIterator) Next() int {
    val := this.data[0]
	this.data = this.data[1:]
	return val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.data) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

