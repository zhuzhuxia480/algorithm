/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
package main

import "strings"
func simplifyPath(path string) string {
	if path == "/" {
		return path
	}
    dirs := strings.Split(path, "/")
	res := make([]string, 0)
	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		} else if dir == ".." {
			if len(res) >=1 && res[len(res)-1] != "/" {
				res  = res[:len(res)-1]
			}			
		} else {
			res = append(res, dir)
		}
	}
	return "/" + strings.Join(res, "/")

}
// @lc code=end

