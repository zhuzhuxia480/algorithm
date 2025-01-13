/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
package main
func longestConsecutive(nums []int) int {
    nMap := make(map[int]bool)
	res := 0
	for _, v := range nums {
		nMap[v] = true		
	}

	for   v := range nMap {
		if !nMap[v-1] {
			tmpSum := 1
			v++
			for nMap[v] {
				tmpSum++
				v++
			}
			if res < tmpSum {
				res = tmpSum
			}
		}
	}
	return res
}
// @lc code=end

