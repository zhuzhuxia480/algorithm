/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
package main
func canFinish(numCourses int, prerequisites [][]int) bool {
    indgree := make([]int, numCourses)
	graph := [][]int{}
	for i := 0; i < numCourses; i++ {
		graph = append(graph, []int{})
	}
	for i := 0; i < len(prerequisites); i++ {
		indgree[prerequisites[i][0]]++
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}
	result := []int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)
		for i := 0; i < len(graph[cur]); i++ {
			indgree[graph[cur][i]]--
			if indgree[graph[cur][i]] == 0 {
				queue = append(queue, graph[cur][i])
			}			
		}
	}
	return len(result) == numCourses
}
// @lc code=end

