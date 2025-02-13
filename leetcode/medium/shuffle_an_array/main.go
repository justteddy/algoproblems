package main

import (
	"math/rand"
	"time"
)

// https://leetcode.com/problems/shuffle-an-array/description/

type Solution struct {
	original []int
	shuffled []int
}

func Constructor(nums []int) Solution {
	s := Solution{
		original: nums,
		shuffled: make([]int, 0, len(nums)),
	}

	for _, v := range s.original {
		s.shuffled = append(s.shuffled, v)
	}

	return s
}

func (this *Solution) Reset() []int {
	return this.original
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(this.shuffled), func(i, j int) {
		this.shuffled[i], this.shuffled[j] = this.shuffled[j], this.shuffled[i]
	})

	return this.shuffled
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
