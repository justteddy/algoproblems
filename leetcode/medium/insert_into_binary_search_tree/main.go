package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 12,
				Left: &TreeNode{
					Val: 15,
				},
			},
		},
	}
	fmt.Println(insertIntoBST(tree, 16))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root
	for {
		if val < curr.Val {
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				break
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				break
			}
			curr = curr.Right
		}
	}

	return root
}
