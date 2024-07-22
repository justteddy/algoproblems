package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	df, result := diff(root.Val, target), root.Val

	var traverse func(n *TreeNode)
	traverse = func(n *TreeNode) {
		if n == nil {
			return
		}

		// edge case - totally equal
		if float64(n.Val) == target {
			result = n.Val
			return
		}

		newDf := diff(n.Val, target)
		if newDf == df && n.Val < result {
			df = newDf
			result = n.Val
		} else if newDf < df {
			df = newDf
			result = n.Val
		}

		if float64(n.Val) > target {
			traverse(n.Left)
		} else {
			traverse(n.Right)
		}
	}

	traverse(root)
	return result
}

func diff(n int, target float64) float64 {
	if target >= 0 {
		return abs(float64(n) - target)
	}

	return float64(n) - target
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
