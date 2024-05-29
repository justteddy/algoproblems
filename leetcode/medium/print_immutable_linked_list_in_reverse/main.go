package main

import "fmt"

type ImmutableListNode struct {
	val  int
	next *ImmutableListNode
}

func (i *ImmutableListNode) getNext() *ImmutableListNode {
	return i.next
}

func (i *ImmutableListNode) printValue() {
	fmt.Println(i.val)
}

func main() {
	node5 := &ImmutableListNode{val: 5, next: nil}
	node4 := &ImmutableListNode{val: 4, next: node5}
	node3 := &ImmutableListNode{val: 3, next: node4}
	node2 := &ImmutableListNode{val: 2, next: node3}
	node1 := &ImmutableListNode{val: 1, next: node2}

	printLinkedListInReverse(node1)
	fmt.Println()
	printLinkedListInReverse_(node1)
	fmt.Println()
}

// recursive solution
func printLinkedListInReverse(head *ImmutableListNode) {
	if head == nil {
		return
	}

	printLinkedListInReverse(head.getNext())
	head.printValue()
}

// stack solution
func printLinkedListInReverse_(head *ImmutableListNode) {
	list := make([]func(), 0)

	next := head
	for next != nil {
		list = append(list, next.printValue)
		next = next.getNext()
	}

	for i := len(list) - 1; i >= 0; i-- {
		list[i]()
	}
}
