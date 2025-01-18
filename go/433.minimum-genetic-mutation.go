/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
package main

import "fmt"
func minMutation(startGene string, endGene string, bank []string) int {
    if startGene == endGene {
		return 0
	}

	bankMap := make(map[string]bool)
	for i := 0; i < len(bank); i++ {
		bankMap[bank[i]] = true
	}

	gene := []byte{'A', 'C', 'G', 'T'}

	if !bankMap[endGene] {
		return -1
	}

	vis := make(map[string]bool)
	vis[startGene] = true

	q := []string{startGene}
	step := 0
	for len(q) > 0 {
		step++
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			curGene := q[i]
			for j := 0; j < 8; j++ {
				for k := 0; k < 4; k++ {
					if gene[k] != curGene[j] {
						newGene := curGene[:j] + string(gene[k]) + curGene[j+1:]
						if newGene == endGene {
							return step
						}
						if !bankMap[newGene] {
							continue
						}
						if !vis[newGene] {
							q = append(q, newGene)
							vis[newGene] = true
						}
					}
				}
			}
		}
		q = q[curLen:]
	}
	return -1
}
// @lc code=end

