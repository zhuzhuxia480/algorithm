/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
package main

import (
	"sort"
)
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
    sort.Ints(nums)
	numLen := len(nums)
	for i := 0; i < numLen - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ps := i+1
		pe := numLen-1
		for ps < pe {
			if nums[ps] + nums[pe] + nums[i] == 0 {
				tmp := []int{nums[i], nums[ps], nums[pe]}
				ret = append(ret, tmp)
				ps++
				for nums[ps] == nums[ps-1] && ps < pe {
					ps++
				}
			} else if nums[ps] + nums[pe] + nums[i] > 0 {
				pe--
			} else if nums[ps] + nums[pe] + nums[i] < 0 {
				ps++
			}
		}
	}
	return ret
}
// @lc code=end

