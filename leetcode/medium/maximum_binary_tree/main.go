package main

import "fmt"

// https://leetcode.com/problems/maximum-binary-tree/
func main() {
	fmt.Println(constructMaximumBinaryTree([]int{6, 2, 4, 1, 8, 7}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIdx := findMaxIdx(nums)
	node := &TreeNode{Val: nums[maxIdx]}
	node.Left = constructMaximumBinaryTree(nums[:maxIdx])
	node.Right = constructMaximumBinaryTree(nums[maxIdx+1:])

	return node
}

func findMaxIdx(nums []int) int {
	max, idx := 0, 0
	for i, v := range nums {
		if v > max {
			idx, max = i, v
		}
	}

	return idx
}
