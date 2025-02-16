/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
package main
func trailingZeroes(n int) int {
	ans := 0
    for n > 0 {
		ans += n/5
		n /= 5
	}
	return ans
}
// @lc code=end

