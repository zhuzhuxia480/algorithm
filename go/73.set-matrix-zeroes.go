/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
package main
func setZeroes(matrix [][]int)  {
    rowMap := make(map[int]bool)
	colMap := make(map[int]bool)

	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = true
				colMap[j] = true
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			_, okRow := rowMap[i]
			_, okCol := colMap[j]
			if okCol || okRow {
				matrix[i][j] = 0
			}
		}
	}
}
// @lc code=end

