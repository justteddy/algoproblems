package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/binary-tree-tilt
func findTilt(root *TreeNode) int {
	total := 0
	var traverse func(n *TreeNode) int
	traverse = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		left := traverse(n.Left)
		right := traverse(n.Right)

		total += abs(left - right)
		return n.Val + left + right
	}

	traverse(root)

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
