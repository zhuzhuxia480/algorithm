package main

func checkIndexSame(index int, strs []string) bool {
	if index >= len(strs[0]) {
		return false
	}
	flag := strs[0][index]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) <= index {
			return false
		}
		if flag != strs[i][index]{
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	len := len(strs[0])
	preFixIndex := 0
	for i := 0; i < len; i++ {
		if !checkIndexSame(preFixIndex, strs) {
			break
		}
		preFixIndex++
	}
	return strs[0][0:preFixIndex]
}
