package main

func reverseString(s []byte) {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		s[i], s[sLen-1-i] = s[sLen-1-i], s[i]
	}
}
