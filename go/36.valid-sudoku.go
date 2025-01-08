/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
package main

const sudoLen = 9

func checkRowValid(board [][]byte) bool {
	for i := 0; i < sudoLen; i++ {
		numberMap := make(map[byte]bool)
		for j := 0; j < sudoLen; j++ {
			if board[i][j] != '.' {
				val := board[i][j]
				if _, ok := numberMap[val]; ok {
					return false
				}
				numberMap[val] = true
			}
		}
	}
	return true
}

func checkColValid(board [][]byte) bool {
	for i := 0; i < sudoLen; i++ {
		numberMap := make(map[byte]bool)
		for j := 0; j < sudoLen; j++ {
			if board[j][i] != '.' {
				val := board[j][i]
				if _, ok := numberMap[val]; ok {
					return false
				}
				numberMap[val] = true
			}
		}
	}
	return true
}

func checkNine(board [][]byte, row int, col int) bool {
	numberMap := make(map[byte]bool)
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if board[i][j] != '.' {
				if _, ok := numberMap[board[i][j]]; ok {
					return false
				}
				numberMap[board[i][j]] = true
			}

		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	return checkRowValid(board) && checkColValid(board) &&
		checkNine(board, 0, 0) && checkNine(board, 0, 3) && checkNine(board, 0, 6) &&
		checkNine(board, 3, 0) && checkNine(board, 3, 3) && checkNine(board, 3, 6) &&
		checkNine(board, 6, 0) && checkNine(board, 6, 3) && checkNine(board, 6, 6)

}

// @lc code=end
