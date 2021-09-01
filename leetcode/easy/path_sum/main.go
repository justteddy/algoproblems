package main

import "fmt"

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
	fmt.Println(hasPathSum(tree, 6))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return pathSum(root, 0, targetSum)

}

func pathSum(node *TreeNode, curr, target int) bool {
	if node == nil {
		return false
	}

	curr += node.Val
	if node.Left == nil && node.Right == nil {
		if curr == target {
			return true
		}
		return false
	}

	left := pathSum(node.Left, curr, target)
	right := pathSum(node.Right, curr, target)

	return left || right
}
