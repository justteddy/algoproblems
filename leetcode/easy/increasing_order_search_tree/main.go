package main

import (
	"fmt"
)

/*
	Given the root of a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only one right child.
	Example 1:

	Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
	Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

	Example 2:
	Input: root = [5,1,7]
	Output: [1,null,5,null,7]

	Constraints:
	The number of nodes in the given tree will be in the range [1, 100].
	0 <= Node.val <= 1000
*/

func main() {
	tree := &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  16,
				Left: nil,
				Right: &TreeNode{
					Val:   18,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 32,
			Left: &TreeNode{
				Val:   25,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  35,
				Left: nil,
				Right: &TreeNode{
					Val:   40,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(increasingBST(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := new(TreeNode)
	curr := dummyNode

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		right := &TreeNode{Val: node.Val}
		curr.Right = right
		curr = right
		traverse(node.Right)
	}

	traverse(root)

	return dummyNode.Right
}
