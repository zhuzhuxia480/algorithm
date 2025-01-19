/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
package main

import "strconv"

var ans []string
var digitMap = [][]byte{
	{},
	{},
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func dfsNum(index int, digits string, tmpans string) {
	if index == len(digits) {
		ans = append(ans, tmpans)
		return
	}
	ni, _:= strconv.Atoi(string(digits[index]))
	for j := 0; j < len(digitMap[ni]); j++ {
		curAns := tmpans + string(digitMap[ni][j])
		dfsNum(index+1, digits, curAns)
	}
	
}

func letterCombinations(digits string) []string {
	ans = []string{}
	if len(digits) == 0 {
		return ans
	}
	dfsNum(0, digits, "")
	return ans
}

// @lc code=end
