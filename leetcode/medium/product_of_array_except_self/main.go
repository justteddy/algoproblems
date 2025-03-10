package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3}))
}

func productExceptSelf(nums []int) []int {
	mp, zeros := 1, 0
	for _, v := range nums {
		if v == 0 {
			zeros++
			continue
		}
		mp *= v
	}

	if zeros > 1 {
		return make([]int, len(nums))
	}

	zeroFound := zeros == 1
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		if v == 0 {
			res = append(res, mp)
		} else {
			if zeroFound {
				res = append(res, 0)
			} else {
				res = append(res, mp/v)
			}
		}
	}

	return res
}
