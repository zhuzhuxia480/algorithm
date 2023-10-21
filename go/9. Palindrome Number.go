package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strLen := len(str)
	for i := 0; i < strLen/2; i++ {
		if str[i] != str[strLen-i-1] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome(1212123))
}
