/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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
func roundToDecimalPlaces(num float64, decimalPlaces int) float64 {
	// 计算 10^decimalPlaces
	scale := math.Pow(10, float64(decimalPlaces))
	// 四舍五入到指定的小数位数
	return math.Round(num*scale) / scale
}
func averageOfLevels(root *TreeNode) []float64 {
    res := []float64{}
	queue := []TreeNode{*root}
	for len(queue) > 0 {
		currentLen := len(queue)
		sum := float64(0)
		for i := 0; i < currentLen; i++ {
			tmpNode := queue[0]
			queue = queue[1:]
			if tmpNode.Left != nil {
				queue = append(queue, *tmpNode.Left)
			}
			if tmpNode.Right != nil {
				queue = append(queue, *tmpNode.Right)
			}
			sum += float64(tmpNode.Val)
		}
		tmpRes := sum / float64(currentLen)
		res = append(res, roundToDecimalPlaces(tmpRes, 5))
	}
	return res
}
// @lc code=end

