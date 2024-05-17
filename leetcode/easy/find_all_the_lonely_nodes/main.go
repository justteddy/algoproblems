package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(getLonelyNodes(treeExample()))
}

func getLonelyNodes(root *TreeNode) []int {
	return dfs(root)
}

func dfs(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(n *TreeNode)
	traverse = func(n *TreeNode) {
		if n == nil {
			return
		}

		if n.Left == nil && n.Right != nil {
			result = append(result, n.Right.Val)
		} else if n.Right == nil && n.Left != nil {
			result = append(result, n.Left.Val)
		}

		traverse(n.Left)
		traverse(n.Right)
	}

	traverse(root)
	return result
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
