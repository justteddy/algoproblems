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

	var traverseLevel func([]*TreeNode) ([]int, []*TreeNode)
	traverseLevel = func(nodes []*TreeNode) ([]int, []*TreeNode) {
		next := make([]*TreeNode, 0)
		vals := make([]int, 0, len(nodes))
		for _, node := range nodes {
			vals = append(vals, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		return vals, next
	}

	result := make([][]int, 0)
	vals := make([]int, 0)
	next := []*TreeNode{root}
	for len(next) != 0 {
		vals, next = traverseLevel(next)
		result = append(result, vals)
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
