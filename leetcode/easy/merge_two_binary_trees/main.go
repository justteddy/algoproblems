package main

import (
	"fmt"
)

/*
	You are given two binary trees root1 and root2.
	Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.
	Return the merged tree.
	Note: The merging process must start from the root nodes of both trees.

	Example 1:
	Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
	Output: [3,4,5,5,4,null,7]

	Example 2:
	Input: root1 = [1], root2 = [1,2]
	Output: [2,2]

	Constraints:
	The number of nodes in both trees is in the range [0, 2000].
	-104 <= Node.val <= 104
*/

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(mergeTrees(tree1, tree2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	newNode := &TreeNode{Val: n1.Val + n2.Val}
	newNode.Left = mergeTrees(n1.Left, n2.Left)
	newNode.Right = mergeTrees(n1.Right, n2.Right)

	return newNode
}
