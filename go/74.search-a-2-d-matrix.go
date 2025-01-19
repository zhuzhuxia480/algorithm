/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
package main

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])

	i, j := 0, row*col-1
	for i <= j {
		mid := (i + j) / 2
		x, y := (mid)/col, mid%col
		if matrix[x][y] < target {
			i = mid + 1
		} else if matrix[x][y] > target {
			j = mid - 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
