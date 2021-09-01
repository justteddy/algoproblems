package main

import (
	"fmt"
)

/*
	Given the root of a binary tree, return its maximum depth.
	A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

	Example 1:
	Input: root = [3,9,20,null,null,15,7]
	Output: 3

	Example 2:
	Input: root = [1,null,2]
	Output: 2

	Example 3:
	Input: root = []
	Output: 0

	Example 4:
	Input: root = [0]
	Output: 1

	Constraints:
	The number of nodes in the tree is in the range [0, 104].
	-100 <= Node.val <= 100
*/

func main() {
	tree := &TreeNode{
		Left: &TreeNode{},
		Right: &TreeNode{
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
	}

	fmt.Println(maxDepth(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
func maxDepth(root *TreeNode) int {
	var depthFn func(node *TreeNode, depth int) int
	depthFn = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth
		}

		left := depthFn(node.Left, depth+1)
		right := depthFn(node.Right, depth+1)

		if left > right {
			return left
		}

		return right

	}

	return depthFn(root, 0)
}

// iterative, BFS solutions
func _maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	nodes := []*TreeNode{root}

	for len(nodes) != 0 {
		next := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		nodes = next
		depth++
	}

	return depth
}
