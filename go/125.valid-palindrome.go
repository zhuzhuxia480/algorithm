/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
package main

import "strings"
func alphanNumeric(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
func isPalindrome(s string) bool {
	sLen := len(s)
	i := 0
	j := sLen - 1
	lowerS := strings.ToLower(s)
	for i < j {
		alphan := alphanNumeric(lowerS[i])
		for !alphan && i < j{
			i++
			alphan = alphanNumeric(lowerS[i])
		}
		alphan = alphanNumeric(lowerS[j])
		for !alphan && i < j{
			j--
			alphan = alphanNumeric(lowerS[j])
		}
		if lowerS[i] != lowerS[j] {
			return false
		}
		i++
		j--
	}
	return true

}

// @lc code=end


