/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start

package main
var ans [][]int
var vis [21]int
func dfsCombine(n int, k int, count int, tmpAns []int) {
	if count == k {
		ans = append(ans, append([]int{}, tmpAns...))
		return 
	}
	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			if len(tmpAns) > 0 && tmpAns[len(tmpAns)-1] > i {
				continue
			}
			vis[i] = 1
			dfsCombine(n, k, count+1, append(tmpAns, i))
			vis[i] = 0
		}
	}
}
func combine(n int, k int) [][]int {
    for i := 1; i <= n; i++ {
		vis[i] = 0
	}
	ans = [][]int{}
	dfsCombine(n, k, 0, []int{})
	return ans
}
// @lc code=end

