package main

import (
	"fmt"
)

/*
	Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

	Example 1:
	Input: root = [3,9,20,null,null,15,7]
	Output: [[3],[9,20],[15,7]]

	Example 2:
	Input: root = [1]
	Output: [[1]]

	Example 3:
	Input: root = []
	Output: []

	Constraints:
	The number of nodes in the tree is in the range [0, 2000].
	-1000 <= Node.val <= 1000
*/
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
	fmt.Println(levelOrder(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		vals := make([]int, 0, len(nodes))
		next := make([]*TreeNode, 0)
		for _, node := range nodes {
			vals = append(vals, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		result = append(result, vals)
		nodes = next
	}

	return result
}

type Queue []*TreeNode

func (q *Queue) Len() int            { return len(*q) }
func (q *Queue) IsEmpty() bool       { return len(*q) == 0 }
func (q *Queue) Enqueue(n *TreeNode) { *q = append(*q, n) }
func (q *Queue) Dequeue() *TreeNode {
	if len(*q) == 0 {
		return nil
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element
	}
}
