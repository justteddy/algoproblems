package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/
func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	stack := make([]string, 0)
	for _, char := range s {
		sym := string(char)
		switch sym {
		case "(":
			stack = append(stack, ")")
		case "[":
			stack = append(stack, "]")
		case "{":
			stack = append(stack, "}")
		default:
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != sym {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
