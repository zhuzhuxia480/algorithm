/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
package main

func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen)
	dp[0] = 1
	ret := 1
	for i := 1; i < numsLen; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}

// @lc code=end
