package main

import "fmt"

/*
	Given the root of a binary tree, return the inorder traversal of its nodes' values.

	Example 1:
	Input: root = [1,null,2,3]
	Output: [1,3,2]

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
	fmt.Println(inorderTraversalIterative(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
func inorderTraversal(root *TreeNode) []int {
	return traverse(root, []int{})
}

func traverse(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = traverse(node.Left, result)
	result = append(result, node.Val)
	result = traverse(node.Right, result)

	return result
}

// iterative solution
func inorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := new(Stack)
	curr := root

	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}

		curr = stack.Pop()
		result = append(result, curr.Val)
		curr = curr.Right
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
