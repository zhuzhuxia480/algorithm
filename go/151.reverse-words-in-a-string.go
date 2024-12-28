/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	res := ""
	firstSpace := false
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "" {
			continue
		}
		if firstSpace == false {
			firstSpace = true
			res += words[i]
		} else {
			res += " " + words[i]
		}
		
	}
	return res
}
// @lc code=end

