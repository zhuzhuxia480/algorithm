/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
package main
func dfsLands(grid [][]byte, flagMap [][]int, i ,j int) {
	row := len(grid)
	col := len(grid[0])
	if i < 0 || i >= row || j < 0 || j >= col {
		return
	}
	if grid[i][j] == '0' || flagMap[i][j] == 1 {
		return
	}

	flagMap[i][j] = 1
	dfsLands(grid, flagMap, i-1,j)
	dfsLands(grid, flagMap, i+1,j)
	dfsLands(grid, flagMap, i,j-1)
	dfsLands(grid, flagMap, i,j+1)
}
func numIslands(grid [][]byte) int {
	res := 0
    row := len(grid)
	col := len(grid[0])
	flagMap := make([][]int, row)
	for i := 0; i < row; i++ {
		flagMap[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' && flagMap[i][j] == 0 {
				dfsLands(grid, flagMap, i, j)
				res++
			}
		}
	}
	return res
}
// @lc code=end

