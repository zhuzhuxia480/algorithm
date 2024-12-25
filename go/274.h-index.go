/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	numsLen := len(citations)
	res := 0
	for i := 1; i <= numsLen; i++ {
		sum := 0
		for j := 0; j < numsLen; j++ {
			if citations[j] >= i {
				sum++
			}
		}
		if sum >= i {
			res = i
		}
	}
	return res
}

// @lc code=end

