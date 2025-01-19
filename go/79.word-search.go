/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
package main
var dirx = []int{-1, 0, 1, 0}
var diry = []int{0, 1, 0, -1}

func clearVis(vis map[int]map[int]bool, row, col int) {
	for i := 0; i < row; i++ {
		vis[i] = make(map[int]bool)
		for j := 0; j < col; j++ {
			vis[i][j] = false
		}
	}
}

func exist(board [][]byte, word string) bool {
    vis := make(map[int]map[int]bool)
	row := len(board)
	col := len(board[0])

	
	
	var dfsFunc func(board [][]byte, index int, x, y int) bool
	ret := false
	dfsFunc = func(board [][]byte, index int, x, y int) bool {
		if index == len(word) {
			ret = true
			return true
		}
	    vis[x][y] = true
		for i := 0; i < len(dirx); i++ {
			nexti := x + dirx[i]
			nextj := y + diry[i]
			if nexti >= 0 && nexti < row && nextj >= 0 && nextj < col && board[nexti][nextj] == word[index] {
				if !vis[nexti][nextj] {
					dfsFunc(board, index+1, nexti, nextj)
					vis[nexti][nextj] = false
				}
			}
		}
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				clearVis(vis, row, col)
				dfsFunc(board, 1, i, j)
				if ret {
					return ret
				}
			}
		}
	}
	return false
}

// @lc code=end


