/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
package main
type iterNode struct{
	key string
	val float64
}
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    graph := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		start := equations[i][0]
		end := equations[i][1]
		val := values[i]
		if graph[start] == nil {
			graph[start] = make(map[string]float64)
		}
		if graph[end] == nil {
			graph[end] = make(map[string]float64, 0)
		}
		graph[start][end] = val
		graph[end][start] = 1/val
	}
	result := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		result[i] = -1.0
		start := queries[i][0]
		end := queries[i][1]
		if graph[start] == nil || graph[end] == nil {
			continue
		}
		if start == end {
			result[i] = 1.0
			continue
		}

		queue := []iterNode{{start, 1}}
		visited := make(map[string]bool)
		visited[start] = true
		for len(queue) > 0 {
			curNode := queue[0]
			if curNode.key == end {
				result[i] = curNode.val
				break
			}
			for key, val := range graph[curNode.key] {
				nexNode := iterNode{
					key: key,
					val: val*curNode.val,
				}
				if visited[key] == false{
					visited[key] = true
					queue = append(queue, nexNode)
				}
			}
			queue = queue[1:]
		}
	}
	return result
}
// @lc code=end

