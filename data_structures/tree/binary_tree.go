package tree

import "leetcode/data_structures/queue"

type BinaryNode struct {
	Val   int
	Left  *BinaryNode
	Right *BinaryNode
}

// TraverseBFS breadth first search traversal
func TraverseBFS(root *BinaryNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	q := new(queue.Queue)
	q.Enqueue(root)

	for !q.IsEmpty() {
		node := q.Dequeue().(*BinaryNode)

		result = append(result, node.Val)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}

	return result
}

// TraverseDFSPreOrderRecursive depth first search traversal preorder
func TraverseDFSPreOrderRecursive(root *BinaryNode) []int {
	result := make([]int, 0)

	var traverse func(*BinaryNode)
	traverse = func(node *BinaryNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// TraverseDFSInOrderRecursive depth first search traversal inorder
func TraverseDFSInOrderRecursive(root *BinaryNode) []int {
	result := make([]int, 0)

	var traverse func(*BinaryNode)
	traverse = func(node *BinaryNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// TraverseDFSPostOrderRecursive depth first search traversal postorder
func TraverseDFSPostOrderRecursive(root *BinaryNode) []int {
	result := make([]int, 0)

	var traverse func(*BinaryNode)
	traverse = func(node *BinaryNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}

	traverse(root)
	return result
}
