/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	gasLen := len(gas)
	i := 0
	for i < gasLen {
		flag, nextStep := checkStep(i, gas, cost, gasLen)
		if flag == true {
			return i
		} else {
			if nextStep+1 <= i {
				return -1
			}
			i = nextStep+1
		}
	}
	return -1
}
func checkStep(start int, gas[]int, cost []int, length int) (bool, int) {
	sum := 0
	for i := 0; i < length; i++ {
		sum += gas[start] - cost[start]
		if sum < 0 {
			return false, start
		}
		start = (start+1) % length
	}
	return true, start
}
// @lc code=end

