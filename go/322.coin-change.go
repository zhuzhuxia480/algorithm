/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
package main
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount+1
	}

	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if coins[i] <= j {
				dp[j] = min(dp[j], dp[j-coins[i]] + 1)
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
	//dp[j] = min(dp[j-1], dp[j-coins[i]] + 1)
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

