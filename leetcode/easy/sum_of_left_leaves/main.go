package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(sumOfLeftLeaves(treeExample()))
	fmt.Println(sumOfLeftLeaves2(treeExample()))
}

// dfs recursive solution
func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	var traverse func(n *TreeNode, isLeft bool)
	traverse = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil { // it is a leaf
			if isLeft { // it is a left leaf
				result += node.Val
			}
			return
		}
		traverse(node.Left, true)
		traverse(node.Right, false)
	}

	traverse(root, false)

	return result
}

func sumOfLeftLeaves2(root *TreeNode) int {
	return dfs(root, false)
}

func dfs(n *TreeNode, isLeft bool) int {
	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil && isLeft {
		return n.Val
	}

	return dfs(n.Left, true) + dfs(n.Right, false)
}

/*
		 [20]
		/    \
     [12]    [32]
     /  \    /   \
  [6]  [16] [25]  [35]
         \           \
         [18]        [40]
*/
func treeExample() *TreeNode {
	return &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  16,
				Left: nil,
				Right: &TreeNode{
					Val:   18,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 32,
			Left: &TreeNode{
				Val:   25,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  35,
				Left: nil,
				Right: &TreeNode{
					Val:   40,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
}
