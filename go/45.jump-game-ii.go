/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
    numsLen := len(nums)
    addFlag := make([]int, numsLen)
	const invalidVal = 999999
	for i := 1; i < numsLen; i++ {
		addFlag[i] = invalidVal
	}

	if numsLen == 1 {
		return 0
	}
	queue := []int{0}
	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]
		currentVal := nums[index]
		for i := 1; i <= currentVal; i++ {
			nextIdx := index + i
			if nextIdx >= numsLen {
				break
			}
			if addFlag[nextIdx] == invalidVal {
				queue = append(queue, nextIdx)
			}
			tmpStep := addFlag[index] + 1
			if addFlag[nextIdx] > tmpStep{
				addFlag[nextIdx] = tmpStep
			}
			
		} 
	}
	return addFlag[numsLen-1]
}
// @lc code=end

