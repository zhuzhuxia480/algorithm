/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
package main
func rob(nums []int) int {
    numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	dp := make([]int, numsLen)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < numsLen; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	ans := 0
	for i := 0; i < numsLen; i++ {
		if ans < dp[i] {
			ans = dp[i]
		}
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

