package main

import "fmt"

// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(sumRootToLeaf(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	result := 0
	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, num int) {
		if node == nil {
			return
		}

		num = (num << 1) | node.Val
		if node.Left == nil && node.Right == nil {
			result += num
			return
		}

		traverse(node.Left, num)
		traverse(node.Right, num)
	}

	traverse(root, 0)
	return result
}
