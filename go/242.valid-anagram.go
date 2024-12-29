/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
package main
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		v, ok := sMap[t[i]]
		if !ok {
			return false
		} else {
			if v <= 0 {
				return false
			}
			sMap[t[i]]--
		}
	}
	return true
}
// @lc code=end

