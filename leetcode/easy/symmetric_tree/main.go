package main

import (
	"fmt"
)

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isSymmetric(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// iterative solution
func isSymmetric(root *TreeNode) bool {
	stack := new(Stack)

	stack.Push(root.Right)
	stack.Push(root.Left)

	for !stack.IsEmpty() {
		n1, n2 := stack.Pop(), stack.Pop()
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		stack.Push(n1.Right)
		stack.Push(n2.Left)

		stack.Push(n1.Left)
		stack.Push(n2.Right)
	}

	return true
}

type Stack []*TreeNode

func (s *Stack) IsEmpty() bool    { return len(*s) == 0 }
func (s *Stack) Push(n *TreeNode) { *s = append(*s, n) }
func (s *Stack) Pop() *TreeNode {
	if len(*s) == 0 {
		return nil
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

// recursive solution
func isSymmetric_(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}

	return isMirror(n1.Left, n2.Right) && isMirror(n1.Right, n2.Left)
}
