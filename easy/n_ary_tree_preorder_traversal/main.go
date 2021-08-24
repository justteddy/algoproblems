package main

import "fmt"

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
	fmt.Println(preorder(tree))
}

type Node struct {
	Val      int
	Children []*Node
}

// recursive solution
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, child := range node.Children {
			traverse(child)
		}
	}

	traverse(root)

	return result
}

// iterative solution
func _preorder(root *Node) []int {
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
			if len(node.Children) != 0 {
				for i := len(node.Children) - 1; i >= 0; i-- {
					stack.Push(node.Children[i])
				}
			}
		}
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
