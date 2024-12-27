/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
package main
func romanToInt(s string) int {
    roman := make(map[string]int)
	roman["I"] = 1
	roman["IV"] = 4
	roman["V"] = 5
	roman["IX"] = 9
	roman["X"] = 10
	roman["XL"] = 40
	roman["L"] = 50
	roman["XC"] = 90
	roman["C"] = 100
	roman["CD"] = 400
	roman["D"] = 500
	roman["CM"] = 900
	roman["M"] = 1000

	sLen := len(s)
	index := 0
	res := 0
	for index < sLen {
		if index + 1 < sLen {
			if v, ok := roman[s[index:index+2]]; ok {
				res += v
				index += 2
				continue
			}
		}
		v, _ := roman[s[index:index+1]]
		res += v
		index++
	}
	return res
}
// @lc code=end

