/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
package main

func convert(s string, numRows int) string {
    if numRows == 1 {
		return s
	}
	sLen := len(s)
	data := make([][]byte, numRows)

	index := 0
	for index < sLen {
		for i := 0; i < numRows && index < sLen; i, index = i+1, index+1 {
			data[i] = append(data[i], s[index])
		}
		
		for i := numRows - 2; i > 0 && index < sLen; i, index = i-1, index+1 {
			data[i] = append(data[i], s[index])
		}
	}
	res := ""
	for _, v := range data {
		res += string(v)
	}
	return res
}
// @lc code=end

