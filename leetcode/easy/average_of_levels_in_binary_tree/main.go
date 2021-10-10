package main

import "fmt"

// https://leetcode.com/problems/average-of-levels-in-binary-tree/
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
	fmt.Println(averageOfLevels(tree))
}

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)

	q := new(Queue)
	q.Enqueue(root)
	for !q.IsEmpty() {
		var count, sum float64
		levelLen := q.Len()

		for levelLen > 0 {
			node := q.Dequeue()
			sum += float64(node.Val)
			count++

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}

			levelLen--
		}

		result = append(result, sum/count)
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) IsEmpty() bool        { return len(*q) == 0 }
func (q *Queue) Len() int             { return len(*q) }
func (q *Queue) Enqueue(el *TreeNode) { *q = append(*q, el) }
func (q *Queue) Dequeue() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}
