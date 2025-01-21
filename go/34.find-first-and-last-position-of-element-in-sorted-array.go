/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
package main
func searchRange(nums []int, target int) []int {
	
    start, end := -1, -1
	index := -1
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}
	l, r := 0, numsLen-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			l = mid+1
		} else {
			r = mid-1
		}
	}
	if index == -1 {
		return []int{-1,-1}
	}
	l, r = 0, index
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			start = mid
			r = mid-1
		}
		if nums[mid] < target {
			l = mid+1
		}
	}
	l, r = index, numsLen-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			end = mid
			l = mid+1
		}
		if nums[mid] > target {
			r = mid-1
		}
	}
	return []int{start, end}
}
// @lc code=end

