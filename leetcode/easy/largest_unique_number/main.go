package main

func largestUniqueNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	mx := -1
	for k, v := range m {
		if v != 1 {
			continue
		}
		if k > mx {
			mx = k
		}
	}

	return mx
}
