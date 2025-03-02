package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/find-leaves-of-binary-tree/description/
func main() {
	/*
			  1
			 / \
			2   3
		   / \
		  4   5
	*/
	fmt.Println(findLeaves(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
}

func findLeaves(root *TreeNode) [][]int {
	maxK := 0
	m := make(map[int][]int)
	var traverse func(n *TreeNode) int
	traverse = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		height := max(traverse(n.Left), traverse(n.Right))
		m[height] = append(m[height], n.Val)
		if height > maxK {
			maxK = height
		}

		return height + 1
	}

	traverse(root)

	res := make([][]int, len(m))
	for i := maxK; i >= 0; i-- {
		res[i] = m[i]
	}

	return res
}
