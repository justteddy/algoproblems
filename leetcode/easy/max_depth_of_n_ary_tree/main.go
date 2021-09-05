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
	fmt.Println(maxDepth(tree))
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	max := 0

	var traverse func(*Node, int)
	traverse = func(node *Node, depth int) {
		if node == nil {
			return
		}

		depth++
		if len(node.Children) == 0 {
			if depth > max {
				max = depth
			}
			return
		}
		for _, child := range node.Children {
			traverse(child, depth)
		}
	}

	traverse(root, 0)

	return max
}
