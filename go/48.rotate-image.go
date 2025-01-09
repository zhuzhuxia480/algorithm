/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
package main
func rotate(matrix [][]int)  {
    row := len(matrix)
	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < row/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][row-j-1]
			matrix[i][row-j-1] = tmp
		}
	}
}
// @lc code=end

