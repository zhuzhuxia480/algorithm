package main

func isValid(s string) bool {
	var stack []uint8
	length := len(s)
	stack = append(stack, s[0])
	for i := 1; i < length; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if s[i] == ')' && stack[len(stack)-1] == '(' ||
				s[i] == ']' && stack[len(stack)-1] == '[' ||
				s[i] == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	isValid("stewtw")

}
