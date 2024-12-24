func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, tmp)
}