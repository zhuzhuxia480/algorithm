/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
    ans := make([]int, numsLen)
	pre := 1
	suf := 1
	for i := 0; i < numsLen; i++ {
		ans[i] = pre
		pre *= nums[i]
	}

	for i := numsLen-1; i >= 0; i-- {
		ans[i] *= suf
		suf *= nums[i]
	}
	return ans
}
// @lc code=end

