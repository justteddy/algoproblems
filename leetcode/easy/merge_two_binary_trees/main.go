package main

import (
	"fmt"
)

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(mergeTrees(tree1, tree2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	newNode := &TreeNode{Val: n1.Val + n2.Val}
	newNode.Left = mergeTrees(n1.Left, n2.Left)
	newNode.Right = mergeTrees(n1.Right, n2.Right)

	return newNode
}
