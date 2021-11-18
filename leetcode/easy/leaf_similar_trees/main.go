package main

import "fmt"

// https://leetcode.com/problems/leaf-similar-trees/
func main() {
	tree1 := &TreeNode{
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
	fmt.Println(leafSimilar(tree1, tree1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	tree1, tree2 := make([]int, 0), make([]int, 0)
	traverse(root1, &tree1)
	traverse(root2, &tree2)

	if len(tree1) != len(tree2) {
		return false
	}

	for i := 0; i < len(tree1); i++ {
		if tree1[i] != tree2[i] {
			return false
		}
	}

	return true
}

func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*result = append(*result, node.Val)
		return
	}

	traverse(node.Left, result)
	traverse(node.Right, result)
}
