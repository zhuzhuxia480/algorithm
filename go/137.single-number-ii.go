/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
package main
func singleNumber(nums []int) int {
    ans := int32(0)
	for i := 0; i < 32; i++ {
		tmpSum := int32(0)
		for _, v := range nums {
			tmpSum += (int32(v)>>i) & 1
		}
		if tmpSum%3 != 0 {
			ans |= (1<<i)
		}
	}
	return int(ans)
}
// @lc code=end

