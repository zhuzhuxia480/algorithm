/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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

import (
	"fmt"
	"math"
	"sort"
)
func listNode(root *TreeNode, vals []int) []int {
	vals = append(vals, root.Val)
	if root.Left != nil {
		vals = listNode(root.Left, vals)
	}
	if root.Right != nil {
		vals = listNode(root.Right, vals)
	}
	return vals
}

func getMinimumDifference(root *TreeNode) int {
    vals :=[]int{}
	vals = listNode(root, vals)
	sort.Ints(vals)
	// fmt.Println(vals)
	minVal := 9999999
	for i := 1; i < len(vals); i++ {
		minVal = int(math.Min(float64(minVal), math.Abs(float64(vals[i]-vals[i-1]))))
	}
	return minVal
}
// @lc code=end

