/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
package main


func lengthOfLastWord(s string) int {
	sLen := len(s)
	end := sLen
	start := -1
	for i := sLen - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}
	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			start = i
			break
		}
		
	}

	res := end - start
	
	return res
}

// @lc code=end
