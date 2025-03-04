/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
    words := make(map[string]bool)
    for _, v := range wordDict {
        words[v] = true
    }

    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if ok, _ := words[s[j:i]]; ok {
                if dp[j] {
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[len(s)]
}

// @lc code=end
