package main

import "fmt"

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := len(nums) / 2
	node := &TreeNode{Val: nums[middle]}
	node.Left = sortedArrayToBST(nums[0:middle])
	node.Right = sortedArrayToBST(nums[middle+1:])

	return node
}
