/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
package main

import "strings"
func wordPattern(pattern string, s string) bool {
    pMap := make(map[byte]string)
	sMap := make(map[string]byte)
	allStr := strings.Split(s, " ")
	if len(pattern) != len(allStr) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		v, ok := pMap[pattern[i]]
		if !ok {
			pMap[pattern[i]] = allStr[i]
			sv, sok := sMap[allStr[i]]
			if !sok {
				sMap[allStr[i]] = pattern[i]
			} else {
				if sv != pattern[i] {
					return false
				}
			}
		} else {
			if v != allStr[i] {
				return false
			}
		}
	}
	return true
}
// @lc code=end

