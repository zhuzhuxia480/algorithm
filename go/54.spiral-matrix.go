/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
package main
func spiralOrder(matrix [][]int) []int {
    top := 0
	right := len(matrix[0])
	bottom := len(matrix)
	left := -1
	res := make([]int, 0)
	pi, pj := 0, 0
	size := len(matrix[0]) * len(matrix)
	for len(res) < size {
		//right
		for pj < right && len(res) < size {
			res = append(res, matrix[pi][pj])
			pj++
		}
		right--
		pi++
		pj--

		//down
		for pi < bottom && len(res) < size {
			res = append(res, matrix[pi][pj])
			pi++
		}
		pi--
		pj--
		bottom--

		//left
		for pj > left && len(res) < size {
			res = append(res, matrix[pi][pj])
			pj--
		}
		pj++
		pi--
		left++

		//up
		for pi > top && len(res) < size {
			res = append(res, matrix[pi][pj])
			pi--
		}
		pi++
		pj++
		top++

	}
	return res
}
// @lc code=end

