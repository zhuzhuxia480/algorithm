/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
package main

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	sindex := 0
	tindex := 0
	if sLen == 0 {
		return true
	}
	for sindex < sLen && tindex < tLen {
		for tindex < tLen && t[tindex] != s[sindex] {
			tindex++
		}
		if tindex < tLen && s[sindex] == t[tindex] {
			sindex++
			tindex++
		}

	}

	if sindex == sLen {
		return true
	}

	return false
}

// @lc code=end
