package main

import "fmt"

/*
	Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

	Example 1:
	Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	Output: [3,9,20,null,null,15,7]

	Example 2:
	Input: inorder = [-1], postorder = [-1]
	Output: [-1]

	Constraints:
	1 <= inorder.length <= 3000
	postorder.length == inorder.length
	-3000 <= inorder[i], postorder[i] <= 3000
	inorder and postorder consist of unique values.
	Each value of postorder also appears in inorder.
	inorder is guaranteed to be the inorder traversal of the tree.
	postorder is guaranteed to be the postorder traversal of the tree.
*/
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
