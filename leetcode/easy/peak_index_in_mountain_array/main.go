package main

import "fmt"

// https://leetcode.com/problems/peak-index-in-a-mountain-array/
func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
}

// binary search solution
func peakIndexInMountainArray(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		med := (low + high) / 2
		if arr[med] < arr[med+1] {
			low = med + 1
		} else {
			high = med
		}
	}

	return low
}

// simple solution
func peakIndexInMountainArray_(arr []int) int {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		if arr[i+1] < arr[i] {
			return i
		}
		if arr[j-1] < arr[j] {
			return j
		}
	}
	return 0
}
