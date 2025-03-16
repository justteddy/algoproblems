package main

import (
	"math"
	"fmt"
)

func main() {
	// Example usage
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 0}}
	result := maxLevelSum(root)
	fmt.Println(result) // Output: 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	m := make(map[int]int)
	var traverse func(*TreeNode, int)
	traverse = func(n *TreeNode, level int) {
		if n == nil {
			return
		}

		m[level] += n.Val
		traverse(n.Left, level+1)
		traverse(n.Right, level+1)
	}

	traverse(root, 1)

	mx := math.MinInt32
	res := 1
	for level, sum := range m {
		if sum > mx {
			mx = sum
			res = level
		}
		if sum == mx {
			if level < res {
				res = level
			}
		}
	}

	return res
}
