/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
package main

func permute(nums []int) [][]int {
	ans := [][]int{}
	vis := make(map[int]bool)
	tmpAns := []int{}
	var dfsFunc func(nums []int, count int)
    dfsFunc = func(nums []int, count int) {
		if count == len(nums) {
			ans = append(ans, append([]int{}, tmpAns...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !vis[nums[i]] {
				vis[nums[i]] = true
				tmpAns = append(tmpAns, nums[i])
				dfsFunc(nums, count+1)
				vis[nums[i]] = false
				tmpAns = tmpAns[0:len(tmpAns)-1]
			}
		}
	}
	dfsFunc(nums, 0)
	return ans
}
// @lc code=end

