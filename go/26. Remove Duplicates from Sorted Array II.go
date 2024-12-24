func removeDuplicates(nums []int) int {
	invalid := -999999
	numsLen := len(nums)
	compareIdx := 0
	duplicated := 0
	for i := 1; i < numsLen; i++ {
		if nums[i] == nums[compareIdx] && i-compareIdx >= 2 {
			nums[i] = invalid
			duplicated++
		} else if nums[i] != nums[compareIdx] {
			compareIdx = i
		}
	}

	i := 0
	for i < numsLen-duplicated {
		if nums[i] == invalid {
			j := i + 1
			for j < numsLen && nums[j] == invalid {
				j++
			}
			nums[i] = nums[j]
			nums[j] = invalid
		}
		i++
	}
	return numsLen - duplicated
}
