package main

import "fmt"

// https://leetcode.com/problems/univalued-binary-tree/
func main() {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(isUnivalTree(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	isUniValued := true
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val != val {
			isUniValued = false
			return
		}
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)

	return isUniValued
}
