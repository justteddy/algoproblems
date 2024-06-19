package main

import "fmt"

func main() {
	v1 := Constructor([]int{1, 0, 0, 2, 3})
	v2 := Constructor([]int{0, 3, 0, 4, 0})
	fmt.Println(v1.dotProduct(v2))
}

type SparseVector struct {
	nonZero map[int]int
}

func Constructor(nums []int) SparseVector {
	nonZero := make(map[int]int)
	for i, num := range nums {
		if num != 0 {
			nonZero[i] = num
		}
	}

	return SparseVector{nonZero}
}

func (sv *SparseVector) getNonZero() map[int]int {
	return sv.nonZero
}

// Return the dotProduct of two sparse vectors
func (sv *SparseVector) dotProduct(vec SparseVector) int {
	res := 0
	nonZeroEls := vec.getNonZero()
	for idx, val := range sv.nonZero {
		if v, ok := nonZeroEls[idx]; ok {
			res += v * val
		}
	}
	return res
}
