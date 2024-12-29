/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
package main
func canConstruct(ransomNote string, magazine string) bool {
    magMap := make(map[byte]int)
	for _, v := range magazine {
		magMap[byte(v)]++
	}

	for i := 0; i < len(ransomNote); i++ {
		v, ok := magMap[ransomNote[i]]
		if !ok {
			return false
		}
		if v <= 0 {
			return false
		}
		magMap[ransomNote[i]]--
	}
	return true
}
// @lc code=end

