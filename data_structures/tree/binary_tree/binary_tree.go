package binary_tree

import (
	"container/list"

	"leetcode/data_structures/queue"
	"leetcode/data_structures/stack"
)

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

// TraverseDFSPreOrderIterative depth first search traversal preorder
func TraverseDFSPreOrderIterative(root *BinaryNode) []int {
	result := make([]int, 0)
	st := new(stack.Stack)
	st.Push(root)

	for !st.IsEmpty() {
		node := st.Pop().(*BinaryNode)
		result = append(result, node.Val)

		if node.Right != nil {
			st.Push(node.Right)
		}
		if node.Left != nil {
			st.Push(node.Left)
		}
	}

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

// TraverseDFSInOrderIterative depth first search traversal inorder
func TraverseDFSInOrderIterative(root *BinaryNode) []int {
	result := make([]int, 0)
	st := new(stack.Stack)
	node := root

	for node != nil || !st.IsEmpty() {
		for node != nil {
			st.Push(node)
			node = node.Left
		}

		node = st.Pop().(*BinaryNode)
		result = append(result, node.Val)
		node = node.Right
	}

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

// TraverseDFSPostOrderIterative depth first search traversal postorder
func TraverseDFSPostOrderIterative(root *BinaryNode) []int {
	ls := list.New()
	st := new(stack.Stack)
	st.Push(root)

	for !st.IsEmpty() {
		node := st.Pop().(*BinaryNode)
		ls.PushFront(node.Val)

		if node.Left != nil {
			st.Push(node.Left)
		}
		if node.Right != nil {
			st.Push(node.Right)
		}
	}

	result := make([]int, 0, ls.Len())
	curr := ls.Front()
	for curr != nil {
		result = append(result, curr.Value.(int))
		curr = curr.Next()
	}

	return result
}
