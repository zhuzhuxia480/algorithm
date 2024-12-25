/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
    result := 0
	sum := 0
	minvalue := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minvalue {
			minvalue = prices[i]
		} else if prices[i] - minvalue > result {
			result = prices[i] - minvalue
			
			if i + 1 < len(prices) {
				if prices[i] >= prices[i+1] {
					minvalue = prices[i]
					sum += result
					result = 0
				}
			}  else if i == len(prices) -1 {
				sum += result
			}
			
		}

		

	}
	return sum
}
// @lc code=end

