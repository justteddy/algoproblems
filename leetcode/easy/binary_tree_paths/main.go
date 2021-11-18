package main

import "fmt"

// https://leetcode.com/problems/binary-tree-paths/
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
	fmt.Println(binaryTreePaths(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)

	var traverse func(*TreeNode, string)
	traverse = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			result = append(result, fmt.Sprintf("%s->%d", path, node.Val))
			return
		}
		traverse(node.Left, fmt.Sprintf("%s->%d", path, node.Val))
		traverse(node.Right, fmt.Sprintf("%s->%d", path, node.Val))
	}

	traverse(root.Left, fmt.Sprintf("%d", root.Val))
	traverse(root.Right, fmt.Sprintf("%d", root.Val))

	if len(result) == 0 {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	return result
}
