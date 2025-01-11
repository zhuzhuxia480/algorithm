/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := [][]int{}
	ret = append(ret, intervals[0])
	for i := 1 ; i < len(intervals); i++ {
		lastVal := ret[len(ret)-1]
		if intervals[i][0] <= lastVal[1]{
			ret[len(ret)-1][1] = max(intervals[i][1], ret[len(ret)-1][1])
		} else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

