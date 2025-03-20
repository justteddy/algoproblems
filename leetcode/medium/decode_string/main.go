package main

import (
	"strconv"
	"leetcode/data_structures/stack"
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	tmp := make([]byte, 0)
	var st stack.GenericStack[any]
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] != ']':
			st.Push(s[i])
		default:
			v := st.Pop().(byte)
			for v != '[' {
				tmp = append(tmp, v)
				v = st.Pop().(byte)
			}
			nums := make([]byte, 0, 3)
			for st.Len() > 0 {
				n := st.Peek().(byte)
				if n >= '0' && n <= '9' {
					nums = append(nums, n)
					st.Pop()
					continue
				}
				break
			}

			for k, j := 0, len(nums)-1; k < j; k, j = k+1, j-1 {
				nums[k], nums[j] = nums[j], nums[k]
			}

			rep, _ := strconv.Atoi(string(nums))
			tmp = bytes.Repeat(tmp, rep)
			for j := len(tmp) - 1; j >= 0; j-- {
				st.Push(tmp[j])
			}
			tmp = make([]byte, 0)
		}
	}

	res := make([]byte, 0, len(st))
	for _, v := range st {
		res = append(res, v.(byte))
	}

	return string(res)
}

//func decodeStringOLD(s string) string {
//	var st stack
//	tmp := make([]byte, 0)
//	res := make([]byte, 0)
//	for i := 0; i < len(s); i++ {
//		switch {
//		case s[i] >= '1' && s[i] <= '9':
//			num, _ := strconv.Atoi(string(s[i]))
//			st.Push(num)
//		case s[i] == '[':
//			continue
//		case s[i] >= 'a' && s[i] <= 'z':
//			if st.Len() > 0 {
//				tmp = append(tmp, s[i])
//			} else {
//				res = append(res, s[i])
//			}
//		case s[i] == ']':
//			cnt, _ := st.Pop()
//
//			if st.Len() > 0 {
//				tmp = append(tmp, bytes.Repeat(tmp, cnt)...)
//				continue
//			} else {
//				res = append(res, bytes.Repeat(tmp, cnt)...)
//				tmp = make([]byte, 0)
//			}
//		}
//	}
//
//	return string(res)
//}
