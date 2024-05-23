package main

import "fmt"

func main() {
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	return process(arr1, arr2, arr3)
}

func process(a1, a2, a3 []int) []int {
	if len(a1) < len(a2) && len(a1) < len(a3) {
		return run(a1, a2, a3)
	}

	if len(a2) < len(a1) && len(a2) < len(a3) {
		return run(a2, a1, a3)
	}

	if len(a3) < len(a1) && len(a3) < len(a2) {
		return run(a3, a1, a2)
	}

	return run(a1, a2, a3)
}

func run(minLenArr, a2, a3 []int) []int {
	dict := make(map[int]int, len(a2)+len(a3))

	for _, v := range a2 {
		dict[v]++
	}
	for _, v := range a3 {
		dict[v]++
	}

	res := make([]int, 0)
	for _, v := range minLenArr {
		cnt, ok := dict[v]
		if ok && cnt == 2 {
			res = append(res, v)
		}
	}

	return res
}
