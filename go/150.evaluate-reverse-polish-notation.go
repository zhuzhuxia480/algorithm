/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
package main

import "strconv"
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		stackLen := len(stack)
		if v == "+" {
			tmpV := stack[stackLen-1] + stack[stackLen-2]
			stack = stack[:stackLen-1]
			stack[stackLen-2] = tmpV
		} else if v == "-" {
			tmpV := stack[stackLen-2] - stack[stackLen-1]
			stack = stack[:stackLen-1]
			stack[stackLen-2] = tmpV
		} else if v == "*" {
			tmpV := stack[stackLen-2] * stack[stackLen-1]
			stack = stack[:stackLen-1]
			stack[stackLen-2] = tmpV
		} else if v == "/" {
			tmpV := stack[stackLen-2] / stack[stackLen-1]
			stack = stack[:stackLen-1]
			stack[stackLen-2] = tmpV
		} else {
			v, err := strconv.Atoi(v)
			if err == nil {
				stack = append(stack, v)
			}
		}
	}
	return stack[0]
}
// @lc code=end

