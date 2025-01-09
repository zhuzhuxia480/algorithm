/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
package main

var dirX = []int{-1, -1, -1, 0, +1, +1, +1, 0}
var dirY = []int{-1, 0, 1, 1, 1, 0, -1, -1}

func getOneCount(board [][]int, x int, y int) int {
	sum := 0
	for i := 0; i < 8; i++ {
		tmpX := x + dirX[i]
		tmpY := y + dirY[i]
		if tmpX >=0 && tmpX < len(board) &&
		   tmpY >=0 && tmpY < len(board[0]) && board[tmpX][tmpY] == 1 {
			sum++
		}
	}
	return sum
}

func gameOfLife(board [][]int) {
	row := len(board)
	col := len(board[0])
	next := make([][]int, row)
	for i := 0; i < row; i++ {
		next[i] = make([]int, col)
		for j := 0; j < col; j++ {
			next[i][j] = board[i][j]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			oneCount := getOneCount(board, i, j)
			if oneCount < 2 {
				next[i][j]=0
			}
			if oneCount == 3 && board[i][j] == 0{
				next[i][j]=1
			}
			if oneCount > 3 && board[i][j] == 1 {
				next[i][j]=0
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = next[i][j]
		}
	}
}

// @lc code=end
