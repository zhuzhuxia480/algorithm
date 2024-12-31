/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
package main
func AbsInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func containsNearbyDuplicate(nums []int, k int) bool {
    nMap := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		nMap[nums[i]] = append(nMap[nums[i]], i)
	}
	for _, val := range nMap {
		if len(val) >= 2 {
			for i := 1; i < len(val); i++ {
				if AbsInt(val[i] - val[i-1]) <= k {
					return true
				}
			}
		}
	}
	return false
}
// @lc code=end

