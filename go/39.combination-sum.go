/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
package main

import "sort"
func combinationSum(candidates []int, target int) [][]int {
    ret := [][]int{}
	tmpAns := []int{}
	sort.Ints(candidates)
	var dfsFunc func(candidates []int, target, sum int) 
	dfsFunc = func(candidates []int, target, sum int) {
		if sum == target {
			ret = append(ret, append([]int{}, tmpAns...))
			return
		}
		for i := 0; i < len(candidates); i++ {
			if candidates[i] + sum <= target &&
			 (len(tmpAns)==0 ||  len(candidates) > 0 && candidates[i] >= tmpAns[len(tmpAns)-1]) {
				tmpAns = append(tmpAns, candidates[i])
				dfsFunc(candidates, target, candidates[i] + sum)
				tmpAns = tmpAns[:len(tmpAns)-1]
			}
		}
	}
	dfsFunc(candidates, target, 0)
	return ret
}
// @lc code=end

