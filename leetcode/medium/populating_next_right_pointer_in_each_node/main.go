package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
func main() {
	tree := &Node{
		Val:   1,
		Left:  &Node{Val: 2},
		Right: &Node{Val: 3},
		Next:  nil,
	}
	fmt.Println(connect(tree))
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := new(Queue)
	queue.Enqueue(root)
	var level float64
	for !queue.IsEmpty() {
		var prev *Node
		for i := 1; i <= int(math.Pow(2, level)); i++ {
			node := queue.Dequeue()
			node.Next = prev
			prev = node
			if node.Left != nil {
				queue.Enqueue(node.Right)
				queue.Enqueue(node.Left)
			}
		}
		level++
	}

	return root
}

type Queue []*Node

func (q *Queue) IsEmpty() bool    { return len(*q) == 0 }
func (q *Queue) Len() int         { return len(*q) }
func (q *Queue) Enqueue(el *Node) { *q = append(*q, el) }
func (q *Queue) Dequeue() *Node {
	if q.IsEmpty() {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}
