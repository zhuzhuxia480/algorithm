/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
package main
func isIsomorphic(s string, t string) bool {
    sMap := make(map[byte]byte)
	vMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		v, ok := sMap[s[i]]
		if !ok {
			sMap[s[i]] = t[i]
			tv, tok := vMap[t[i]]
			if !tok {
				vMap[t[i]] = s[i]
			} else {
				if tv != s[i] {
					return false
				}
			}
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}
// @lc code=end

