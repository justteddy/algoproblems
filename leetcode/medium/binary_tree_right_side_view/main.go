package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	seen := make([]int, 100)
	for i := range seen {
		seen[i] = -101
	}

	var traverse func(n *TreeNode, depth int)
	traverse = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}

		seen[depth] = n.Val
		traverse(n.Left, depth+1)
		traverse(n.Right, depth+1)
	}

	traverse(root, 0)

	for i, v := range seen {
		if v == -101 {
			return seen[:i]
		}
	}

	return seen
}
