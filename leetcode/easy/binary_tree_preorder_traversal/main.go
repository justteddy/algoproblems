package main

import "fmt"

/*
	Given the root of a binary tree, return the preorder traversal of its nodes' values.

	Example 1:
	Input: root = [1,null,2,3]
	Output: [1,2,3]

	Example 2:
	Input: root = []
	Output: []

	Example 3:
	Input: root = [1]
	Output: [1]

	Example 4:
	Input: root = [1,2]
	Output: [1,2]

	Example 5:
	Input: root = [1,null,2]
	Output: [1,2]

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
	fmt.Println(preorderTraversal(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// iterative solution
func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := new(Stack)
	stack.Push(root)

	for !stack.IsEmpty() {
		node := stack.Pop()
		if node != nil {
			result = append(result, node.Val)
			stack.Push(node.Right)
			stack.Push(node.Left)
		}
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
