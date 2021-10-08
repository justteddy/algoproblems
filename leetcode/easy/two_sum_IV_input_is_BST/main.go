package main

import "fmt"

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
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
	fmt.Println(findTarget(tree, 18))
}

func findTarget(root *TreeNode, k int) bool {
	hm := make(map[int]struct{})

	var found bool
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		if _, ok := hm[k-node.Val]; ok {
			found = true
			return
		}
		hm[node.Val] = struct{}{}
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)

	return found
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
