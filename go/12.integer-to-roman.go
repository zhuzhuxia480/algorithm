/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
package main

func intToRoman(num int) string {
    romans := make(map[int]string)
	romans[1] = "I"
	romans[2] = "II"
	romans[3] = "III"
	romans[4] = "IV"
	romans[5] = "V"
	romans[6] = "VI"
	romans[7] = "VII"
	romans[8] = "VIII"
	romans[9] = "IX"

	romans[10] = "X"
	romans[20] = "XX"
	romans[30] = "XXX"
	romans[40] = "XL"
	romans[50] = "L"
	romans[60] = "LX"
	romans[70] = "LXX"
	romans[80] = "LXXX"
	romans[90] = "XC"

	romans[100] = "C"
	romans[200] = "CC"
	romans[300] = "CCC"
	romans[400] = "CD"
	romans[500] = "D"
	romans[600] = "DC"
	romans[700] = "DCC"
	romans[800] = "DCCC"
	romans[900] = "CM"

	romans[1000] = "M"
	romans[2000] = "MM"
	romans[3000] = "MMM"

	res := ""
	mul := 1000
	for num > 0 {
		quotient := num / mul
		if quotient != 0 {
			res += romans[quotient*mul]
		}
		num %= mul
		mul /= 10
	}
	return res
}
// @lc code=end

