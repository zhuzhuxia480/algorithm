func majorityElement(nums []int) int {
	result := 0
	count := 0
    for _, value := range nums {
		if count == 0 {
			result = value
		}
		if value == result {
			count++
		} else {
			count--
		}
	}
	return result
}