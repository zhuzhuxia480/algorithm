/*
 * @lc app=leetcode id=909 lang=golang
 *
 * [909] Snakes and Ladders
 */

// @lc code=start

package main
func getId(num , n int) (x, y int) {
	x = n - 1 - (num-1)/n
	y = (num-1)%n
	if ((num-1)/n)%2 == 1 {
		y = n - 1 - y
	}
	return x, y
}

type IterNode struct {
	id int
	step int
}
func snakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	visited[1] = true
	queue := []IterNode{{1, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 1; i < 7; i++ {
			nextId := cur.id + i
			if nextId > n*n {
				break
			}
			x, y := getId(nextId, n)
			if board[x][y] != -1 {
				nextId = board[x][y]
			}
			if nextId == n*n {
				return cur.step + 1
			}
			if !visited[nextId] {
				visited[nextId] = true
				queue = append(queue, IterNode{nextId, cur.step+1})
			}
		}
	}
	return -1
}
// @lc code=end

