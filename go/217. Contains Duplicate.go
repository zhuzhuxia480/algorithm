package main

func containsDuplicate(nums []int) bool {
	flag := map[int]bool{}
	for _, num := range nums {
		if flag[num] == true {
			return true
		}
		flag[num] = true
	}
	return false
}
