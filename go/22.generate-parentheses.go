/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
package main

func checkValid(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	q := []byte{}
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			q = append(q, []byte(data)[i])
		} else {
			if len(q) == 0 || q[len(q)-1] != '(' {
				return false
			}
			q = q[:len(q)-1]
		}
	}
	if len(q) == 0 {
		return true
	}
	return false
}
func generateParenthesis(n int) []string {
    ans := []string{}
	tmpAns := []byte{}
	var dfsFunc func(n int, count int)
	dfsFunc = func(n int, count int)  {
		if count == n {
			if checkValid(tmpAns) {
				ans = append(ans, string(tmpAns))
			}
			return
		}
		tmpAns = append(tmpAns, '(')
		dfsFunc(n, count+1)
		tmpAns = tmpAns[:len(tmpAns)-1]
		
		tmpAns = append(tmpAns, ')')
		dfsFunc(n, count+1)
		tmpAns = tmpAns[:len(tmpAns)-1]
	}
	dfsFunc(2*n, 0)
	return ans
}
// @lc code=end

