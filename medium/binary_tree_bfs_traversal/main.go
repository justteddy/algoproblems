package main

import (
	"fmt"
)

/*
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

// recursive solution
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
