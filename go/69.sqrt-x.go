/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
package main
func mySqrt(x int) int {
    if x== 0 || x ==1 {
		return x
	}
	i := 0
	for i = 1; i <= x/2; i++ {
		if i*i == x {
			return i
		}
		if i *i > x {
			break
		}
	}
	
	return i-1
}
// @lc code=end

