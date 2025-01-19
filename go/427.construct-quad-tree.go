/*
 * @lc app=leetcode id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

// @lc code=start
/**
 * Definition for a QuadTree node.*/
 package main
// type Node struct {
//     Val bool
//     IsLeaf bool
//     TopLeft *Node
//     TopRight *Node
//     BottomLeft *Node
//     BottomRight *Node
// }


func checkLeaf(grid [][]int, sx, sy, row, col int) bool {
	flag := grid[sx][sy]
	for i := sx; i < row+sx; i++ {
		for j := sy; j < col+sy; j++ {
			if grid[i][j] != flag {
				return false
			}
		}
	}
	return true
}
func buildQuadTree(grid [][]int, sx, sy, row, col int) *Node {
	if row == 1 && col == 1 {
		return &Node{
			Val: grid[sx][sy] == 1,
			IsLeaf: true,
		}
	}

	if !checkLeaf(grid, sx, sy, row, col) {
		topLeft := buildQuadTree(grid, sx, sy, row/2, col/2)
		topRight := buildQuadTree(grid, sx, sy+col/2, row/2, col/2)
		bottomLeft := buildQuadTree(grid, sx+row/2, sy, row/2, col/2)
		bottomRight := buildQuadTree(grid, sx+row/2, sy+col/2, row/2, col/2)
		return &Node{
			IsLeaf: false,
			TopLeft: topLeft,
			TopRight: topRight,
			BottomLeft: bottomLeft,
			BottomRight: bottomRight,
		}
	} else {
		return &Node{
			Val: grid[sx][sy] == 1,
			IsLeaf: true,
		}
	}
}
func construct(grid [][]int) *Node {
    return buildQuadTree(grid, 0, 0, len(grid), len(grid[0]))
}
// @lc code=end


