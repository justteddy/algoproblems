package main

import (
	"container/list"
	"fmt"
)

/*
	Given the root of a binary tree, return the postorder traversal of its nodes' values.

	Example 1:
	Input: root = [1,null,2,3]
	Output: [3,2,1]

	Example 2:
	Input: root = []
	Output: []

	Example 3:
	Input: root = [1]
	Output: [1]

	Example 4:
	Input: root = [1,2]
	Output: [2,1]

	Example 5:
	Input: root = [1,null,2]
	Output: [2,1]

	Constraints:
	The number of the nodes in the tree is in the range [0, 100].
	-100 <= Node.val <= 100
*/
func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(postorderTraversalIterative(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}
	traverse(root)
	return result
}

// iterative solution
func postorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	list := list.New()
	stack := new(Stack)
	stack.Push(root)

	for !stack.IsEmpty() {
		curr := stack.Pop()
		list.PushFront(curr.Val)
		if curr.Left != nil {
			stack.Push(curr.Left)
		}
		if curr.Right != nil {
			stack.Push(curr.Right)
		}
	}

	result := make([]int, 0, list.Len())
	head := list.Front()
	for head != nil {
		result = append(result, head.Value.(int))
		head = head.Next()
	}

	return result
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
