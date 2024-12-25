/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	numsLen := len(nums)
    addFlag := make([]int, numsLen)
	if numsLen == 1 {
		return true
	}
	queue := []int{0}
	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]
		for i := 1; i <= nums[index]; i++ {
			nextIdx := index + i
			if nextIdx == numsLen - 1 {
				return true
			}
			if addFlag[nextIdx] == 0 {
				queue = append(queue, nextIdx)
				addFlag[nextIdx] = 1
			}
		} 
	}
	return false
}
// @lc code=end

