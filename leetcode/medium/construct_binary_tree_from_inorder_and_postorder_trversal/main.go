package main

import "fmt"

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	for i, v := range inorder {
		if v == root.Val {
			leftCount := len(inorder[0:i])
			leftPart := postorder[0:leftCount]
			rightPart := postorder[leftCount : len(postorder)-1]

			root.Left = buildTree(inorder[0:i], leftPart)
			root.Right = buildTree(inorder[i+1:], rightPart)
			break
		}
	}

	return root
}
