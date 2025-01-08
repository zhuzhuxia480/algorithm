/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
package main
func singleNumber(nums []int) int {
    numsLenL := len(nums)
	ret := nums[0]
	for i := 1; i < numsLenL; i++ {
		ret = ret ^ nums[i]
	}
	return ret
}
// @lc code=end

