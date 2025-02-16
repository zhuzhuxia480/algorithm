/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
package main

import "fmt"
func findMin(nums []int) int {
    numsLen := len(nums)
	l, r := 0, numsLen-1
	if nums[numsLen-1] >= nums[0] {
		return nums[0]
	}
	for l <= r {
		if l == r {
			return nums[l]
		}
		mid := (l+r)/2
		fmt.Println(l, r, mid)
		if nums[mid] < nums[0] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[0]
}
// @lc code=end

