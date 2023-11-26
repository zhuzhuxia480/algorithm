package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
        invertTree(root.Right)
	}
	return root
}
