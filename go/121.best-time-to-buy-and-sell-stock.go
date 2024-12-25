/*
//Time limited
func maxProfit(prices []int) int {
    result := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			tmp := prices[j] - prices[i] 
			if tmp > 0 && tmp > result {
				result = tmp
			}
		}
	}
	return result
}
*/
/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
    result := 0
	minvalue := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minvalue {
			minvalue = prices[i]
		} else if prices[i] - minvalue > result {
			result  = prices[i] - minvalue
		}
	}
	return result
}
// @lc code=end

