/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func getIndex(i int, nums[] int) int {
	if i == -1 || i == len(nums) {
		return math.MinInt
	}
	return nums[i]
}
func findPeakElement(nums []int) int {
    numsLen := len(nums)
	
	i, j := 0, numsLen-1
	for i < j {
		mid := (i+j)/2
		midV := getIndex(mid, nums)
		if midV > getIndex(mid-1, nums) && midV > getIndex(mid+1, nums) {
			return mid
		} else if midV > getIndex(mid-1, nums) {
			i = mid+1
		} else {
			j = mid-1
		}
		
	}
	return j
}
// @lc code=end

