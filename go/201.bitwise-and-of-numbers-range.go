/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
package main
func rangeBitwiseAnd(left int, right int) int {
    count := 0
	for left < right {
		right >>= 1
		left >>= 1
		count++
	}
	return right << count
}
// @lc code=end

