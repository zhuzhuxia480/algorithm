/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
package main
func minSubArrayLen(target int, nums []int) int {
    numsLen := len(nums)
	pStart := 0
	sum := 0
	ret := numsLen + 1
	flag := false
	for i := 0; i < numsLen; i++ {
		sum += nums[i]
		if sum >= target {
			flag = true
			for pStart < i && sum-nums[pStart]>=target {
				sum -= nums[pStart]
				pStart++
			}
			tmpLen := i - pStart + 1
			if ret > tmpLen {
				ret = tmpLen
			}
		}
		
	}
	if !flag {
		return 0
	}
	return ret
}
// @lc code=end

