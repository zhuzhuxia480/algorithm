/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
package main

import "strconv"
func summaryRanges(nums []int) []string {
	numsLen := len(nums)
	start := 0
	end := 0
	res := []string{}
	if numsLen == 0 {
		return res
	}
	for i := 1; i < numsLen; i++ {
		if nums[i] > nums[i-1]+1 {
			end = i - 1
			if end == start {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
			}
			start = i
		}
	}
	end = numsLen - 1
	if end == start {
		res = append(res, strconv.Itoa(nums[start]))
	} else {
		res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
	}
	return res
}

// @lc code=end

