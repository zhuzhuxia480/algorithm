/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
package main
func buildMaxHeap(nums []int) {
	length := len(nums)
	for i := length/2-1; i >= 0; i-- {
		moveDown(nums, i)
	}
}

func moveDown(nums []int, index int) {
	length := len(nums)
	left := index*2+1
	right := index*2+2
	largest := index
	if left < length && nums[left] > nums[index] {
		largest = left
	}
	if right < length && nums[right] > nums[largest] {
		largest = right
	}
	if largest != index {
		tmp := nums[index]
		nums[index ] = nums[largest]
		nums[largest] = tmp
		moveDown(nums, largest)
	}
}

func delete(nums *[]int) int {
	ret := (*nums)[0]
	length := len(*nums)
	(*nums)[0] = (*nums)[length-1]
	(*nums) = (*nums)[:length-1]
	moveDown((*nums), 0)
	return ret
}
func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums)
	ret := 0
    for i := 0; i < k; i++ {
		ret = delete(&nums)
	}
	return ret
}
// @lc code=end

