/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
package main
func maxArea(height []int) int {
    res := 0
	i := 0
	j := len(height)-1
	for i < j {
		area := (j-i) * minInt(height[i], height[j])
		if res < area {
			res = area
		}
		if height[i] <= height[j] {
			i++
		} else if height[i] > height[j] {
			j--
		}
	}
	return res
}

func minInt(x int, y int) int {
    if x < y {
        return x
    }
    return y
}
// @lc code=end

