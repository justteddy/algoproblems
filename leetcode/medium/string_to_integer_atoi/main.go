package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

func myAtoi(s string) int {
	leadingZero := false
	var sign *bool
	digits := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || s[i] == '.' { // non digit character
			break
		}
		if s[i] == ' ' { // leading whitespace
			if leadingZero || sign != nil || len(digits) != 0 {
				break
			}
			continue
		}

		if len(digits) == 0 && s[i] == '0' { // leading zeros
			leadingZero = true
			continue
		}

		if s[i] == '+' {
			if leadingZero || sign != nil || len(digits) != 0 {
				break
			}
			sign = ptrBool(true)
			continue
		}

		if s[i] == '-' {
			if leadingZero || sign != nil || len(digits) != 0 {
				break
			}
			sign = ptrBool(false)
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, int(s[i]-'0'))
		}
	}

	if len(digits) == 0 {
		return 0
	}

	res := 0
	for i := 0; i < len(digits); i++ {
		res += digits[i] * pow10(len(digits)-i-1)
		if res < 0 {
			return math.MaxInt32
		}
		if res > math.MaxInt32 || -res < math.MinInt32 {
			if sign != nil && !*sign {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	if sign != nil && !*sign {
		res *= -1
	}

	return res
}

func pow10(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		if res > math.MaxInt32 {
			return res
		}
		res *= 10
	}

	return res
}

func ptrBool(v bool) *bool {
	return &v
}
