package main
func isPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		tmp := int64(i) * int64(i)
		if tmp == int64(num) {
			return true
		}
		if tmp > int64(num) {
			return false
		}
	}
	return false
}
