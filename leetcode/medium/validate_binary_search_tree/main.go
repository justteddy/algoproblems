package main

import "fmt"

// https://leetcode.com/problems/validate-binary-search-tree/
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
	fmt.Println(isValidBST(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Val <= *min) || (max != nil && root.Val >= *max) {
		return false
	}

	return isValid(root.Left, min, &root.Val) && isValid(root.Right, &root.Val, max)
}
