package main

import "fmt"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func main() {
	q := &TreeNode{
		Val: 15,
	}
	p := &TreeNode{
		Val: 12,
		Left: &TreeNode{
			Val: 15,
		},
	}
	tree := &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 12,
				Right: &TreeNode{
					Val: 15,
				},
			},
		},
	}
	fmt.Println(lowestCommonAncestor(tree, p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	min, max := minAndMax(p, q)
	if root.Val >= min.Val && root.Val <= max.Val {
		return root
	}

	if root.Val > max.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return lowestCommonAncestor(root.Right, p, q)
}

func minAndMax(p, q *TreeNode) (*TreeNode, *TreeNode) {
	if p.Val < q.Val {
		return p, q
	}
	return q, p
}
