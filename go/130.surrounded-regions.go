/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
package main

var dir = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func bfs(board [][]byte, i, j int) {
	if board[i][j] != 'O' {
		return
	}

	row := len(board)
	col := len(board[0])
	allNode := [][]int{}
	queue := [][]int{}
	queue = append(queue, []int{i, j})
	allNode = append(allNode, []int{i, j})
	board[i][j] = 'A'
	flag := false
	if i == 0 || i == row-1 || j == 0 || j == col-1 {
		flag = true
	}
	for len(queue) > 0 {
		curNode := queue[0]
		for i := 0; i < len(dir); i++ {
			tmpi := curNode[0] + dir[i][0]
			tmpj := curNode[1] + dir[i][1]
			if tmpi >= 0 && tmpi < row && tmpj >= 0 && tmpj < col && board[tmpi][tmpj] == 'O' {
				queue = append(queue, []int{tmpi, tmpj})
				allNode = append(allNode, []int{tmpi, tmpj})
				board[tmpi][tmpj] = 'A'
				if tmpi == 0 || tmpi == row-1 || tmpj == 0 || tmpj == col-1 {
					flag = true
				}
			}
		}
		queue = queue[1:]
	}
	for _, point := range allNode {
		if flag {
			board[point[0]][point[1]] = 'W'
		} else {
			board[point[0]][point[1]] = 'X'
		}
	}
}

func solve(board [][]byte) {
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				bfs(board, i, j)
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'W' {
				board[i][j] = 'O'
			}
		}
	}
}

// @lc code=end
