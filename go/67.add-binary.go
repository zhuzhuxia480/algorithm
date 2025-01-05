/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
package main

func addBinary(a string, b string) string {
	bytesA := []byte(a)
	bytesB := []byte(b)

	if len(bytesA) < len(bytesB) {
		bytesA, bytesB = bytesB, bytesA
	}

	carry := byte(0)
	i, j := len(bytesA)-1, len(bytesB)-1
	result := make([]byte, len(bytesA)+1)
	rIndex := len(result) - 1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += bytesA[i] - '0'
			i--
		}
		if j >= 0 {
			sum += bytesB[j] - '0'
			j--
		}
		result[rIndex] = sum%2 + '0'
		rIndex--
		carry = sum / 2
	}
	if result[0] == 0 {
		result = result[1:]
	}
	return string(result)
}

// @lc code=end
