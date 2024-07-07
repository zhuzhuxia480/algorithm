package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ret []int
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++
		}
	}
	for i < m {
		ret = append(ret, nums1[i])
		i++
	}
	for j < n {
		ret = append(ret, nums2[j])
		j++
	}

	for i = 0; i < len(ret); i++ {
		nums1[i] = ret[i]
	}
}
