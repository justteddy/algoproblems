package main

import "fmt"

// https://leetcode.com/problems/construct-string-from-binary-tree/
func main() {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(tree2str(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}

	if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	}

	return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
}
