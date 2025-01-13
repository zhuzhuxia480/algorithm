/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(postorder)
    if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{
			Val: inorder[0],
			Left: nil,
			Right: nil,
		}
	}

	
	root := &TreeNode{
		Val: postorder[length-1],
	}
	index := 0
	for i := 0; i < length; i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	root.Left = buildTree(inorder[0:index], postorder[0:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:length-1])
	return root
}
// @lc code=end

