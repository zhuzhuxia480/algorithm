/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
package main

import "sort"
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	merge := [][]int{points[0]}
	for i := 1; i < len(points); i++ {
		lastVal := merge[len(merge)-1]
		if lastVal[1] >= points[i][0] {
			merge[len(merge)-1][0] = max(merge[len(merge)-1][0], points[i][0])
			merge[len(merge)-1][1] = min(merge[len(merge)-1][1], points[i][1])
		} else {
			merge = append(merge, points[i])
		}
	}
	return len(merge)
}
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

