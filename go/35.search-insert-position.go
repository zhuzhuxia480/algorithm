/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
package main

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if nums[start] < target {
		return start + 1
	} else  {
		return start
	}
}

// @lc code=end
