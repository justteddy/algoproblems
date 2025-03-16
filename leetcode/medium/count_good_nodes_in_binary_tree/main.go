package main

// https://leetcode.com/problems/count-good-nodes-in-binary-tree
func main() {
	// Example usage
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	result := goodNodes(root)
	println(result) // Output: 4
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	res := 0
	max := root.Val - 1 // dirty hack
	var traverse func(*TreeNode, int)
	traverse = func(n *TreeNode, max int) {
		if n == nil {
			return
		}

		if n.Val >= max {
			res++
			max = n.Val
		}
		traverse(n.Left, max)
		traverse(n.Right, max)
	}

	traverse(root, max)
	return res
}
