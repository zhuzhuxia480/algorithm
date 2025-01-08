/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// @lc code=end
