/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
package main
func hammingWeight(n int) int {
    sum := 0
	for n > 0 {
		sum += n %2
		n /= 2
	}
	return sum
}
// @lc code=end

