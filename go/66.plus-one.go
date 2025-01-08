/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
package main
func plusOne(digits []int) []int {
    digitsLen := len(digits)
	ret := make([]int, digitsLen+1)
	carry := 0
	sum := 0
	for i := digitsLen - 1; i >= 0; i-- {
		if i == digitsLen - 1 {
			sum = carry + digits[i]+1
		}else {
			sum = carry + digits[i]
		}
		
		ret[i+1] = sum % 10
		carry = sum / 10
	}
	if carry > 0 {
		ret[0] = carry
	}
	if ret[0] == 0 {
		ret = ret[1:]
	}
	return ret
}
// @lc code=end

