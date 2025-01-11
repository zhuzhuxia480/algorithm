/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
package main	
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
func insert(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
	return merge(intervals)
}
// @lc code=end

