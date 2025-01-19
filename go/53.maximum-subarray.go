/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start

package main
func maxSubArray(nums []int) int {
    ans := nums[0]
	preSum := 0
	for _, v := range nums {
		preSum = max(preSum + v, v)
		ans = max(preSum, ans)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

