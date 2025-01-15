/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 package main
func dfs(node * Node, visitMap map[*Node]bool, nodeMap map[*Node]*Node) {
	if node == nil {
		return 
	}
	if _, ok := nodeMap[node]; ok {
		return
	}

	if _, ok := visitMap[node]; ok {
		return
	}
	visitMap[node] = true
	nodeMap[node] = &Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}
	for i := 0; i < len(node.Neighbors); i++ {
		dfs(node.Neighbors[i], visitMap, nodeMap)
		nodeMap[node].Neighbors = append(nodeMap[node].Neighbors, nodeMap[node.Neighbors[i]]) 
	}
}

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil		
	}
	visitMap := make(map[*Node]bool)
	nodeMap := make(map[*Node]*Node)
	dfs(node, visitMap, nodeMap)
	return nodeMap[node]
}
// @lc code=end

