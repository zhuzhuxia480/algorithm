/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
package main
func twoSum(numbers []int, target int) []int {
	numLen := len(numbers)
	i := 0
	j := numLen - 1
	for i < j {
		if numbers[i] + numbers[j] > target {
			j--
		} else if numbers[i] + numbers[j] < target {
			i++
		} else {
			return []int{i+1, j+1}
		}
	}
	return []int{i+1, j+1}
}
// @lc code=end

