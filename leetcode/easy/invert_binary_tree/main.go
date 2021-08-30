package main

import "fmt"

/*
	Given the root of a binary tree, invert the tree, and return its root.

	Example 1:
	Input: root = [4,2,7,1,3,6,9]
	Output: [4,7,2,9,6,3,1]

	Example 2:
	Input: root = [2,1,3]
	Output: [2,3,1]

	Example 3:
	Input: root = []
	Output: []

	Constraints:
	The number of nodes in the tree is in the range [0, 100].
	-100 <= Node.val <= 100
*/
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
	fmt.Println(invertTree(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l, r := root.Left, root.Right

	root.Left = invertTree(r)
	root.Right = invertTree(l)

	return root
}
