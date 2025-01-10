/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
package main
func groupAnagrams(strs []string) [][]string {
    countMap := make(map[[26]int] []string)
	for _, str := range strs {
		tmpCnt := [26]int{}
		for _, c := range str {
			tmpCnt[c-'a']++
		}
		countMap[tmpCnt] = append(countMap[tmpCnt], str)
	}
	res := [][]string{}
	for _, v := range countMap {
		res = append(res, v)
	}
	return res
}
// @lc code=end

