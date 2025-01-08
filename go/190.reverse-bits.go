/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
package main

import "strconv"

func reverseBits(num uint32) uint32 {
	numString := strconv.FormatUint(uint64(num), 2)
	numsByte := []byte(numString)

	for i, j := 0, len(numString)-1; i < j; i, j = i+1, j-1 {
		numsByte[i], numsByte[j] = numsByte[j], numsByte[i]
	}
	for i := len(numsByte); i < 32; i++ {
		numsByte = append(numsByte, '0')
	}

	ret, err := strconv.ParseUint(string(numsByte), 2, 32)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

// @lc code=end
