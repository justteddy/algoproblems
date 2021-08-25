package main

import (
	"container/list"
	"fmt"
)

/*
	Given the root of an n-ary tree, return the postorder traversal of its nodes' values.
	Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)

	Example 1:
	Input: root = [1,null,3,2,4,null,5,6]
	Output: [5,6,3,2,4,1]

	Example 2:
	Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
	Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]

	Constraints:
	The number of nodes in the tree is in the range [0, 104].
	0 <= Node.val <= 104
	The height of the n-ary tree is less than or equal to 1000.

	Follow up: Recursive solution is trivial, could you do it iteratively?
*/

func main() {
	tree := &Node{
		Val: 10,
		Children: []*Node{
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 3,
					},
					{
						Val: 4,
					},
				},
			},
			{
				Val: 8,
			},
		},
	}
	fmt.Println(postorder(tree))
}

type Node struct {
	Val      int
	Children []*Node
}

// recursive solution
func _postorder(root *Node) []int {
	result := make([]int, 0)

	var traverse func(*Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}

		for _, child := range node.Children {
			traverse(child)
		}

		result = append(result, node.Val)
	}

	traverse(root)

	return result
}

// iterative solution
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := new(Stack)
	ls := list.New()

	stack.Push(root)
	for !stack.IsEmpty() {
		node := stack.Pop()
		ls.PushFront(node.Val)

		for _, child := range node.Children {
			stack.Push(child)
		}
	}

	result := make([]int, 0, ls.Len())
	curr := ls.Front()
	for curr != nil {
		result = append(result, curr.Value.(int))
		curr = curr.Next()
	}

	return result
}

type Stack []*Node

func (s *Stack) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack) Push(n *Node)  { *s = append(*s, n) }
func (s *Stack) Pop() *Node {
	if len(*s) == 0 {
		return nil
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}
