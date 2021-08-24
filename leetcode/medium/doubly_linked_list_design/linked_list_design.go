package doubly_linked_list_design

type MyLinkedList struct {
	head   *Node
	length int
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index > l.length-1 {
		return -1
	}

	current := l.head
	for i := 0; current != nil; i++ {
		if i == index {
			return current.val
		}
		current = current.next
	}

	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	newHead := &Node{
		val:  val,
		next: nil,
		prev: nil,
	}

	if l.head != nil {
		newHead.next = l.head
		l.head.prev = newHead
	}

	l.head = newHead
	l.length++
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.head = &Node{
			val:  val,
			next: nil,
			prev: nil,
		}
		l.length++
		return
	}

	current := l.head
	for current != nil {
		if current.next == nil {
			current.next = &Node{
				val:  val,
				next: nil,
				prev: current,
			}
			break
		}
		current = current.next
	}
	l.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.length {
		return
	}

	if index == l.length {
		l.AddAtTail(val)
		return
	}

	if index == 0 {
		l.AddAtHead(val)
		return
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	n := &Node{
		val:  val,
		next: current,
		prev: current.prev,
	}

	current.prev.next = n
	current.prev = n
	l.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index > l.length-1 {
		return
	}

	if index == 0 && l.head != nil {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		l.length--
		return
	}

	current := l.head
	for i := 0; i <= index; i++ {
		if i != index {
			current = current.next
			continue
		}
		current.prev.next = current.next
		if current.next != nil {
			current.next.prev = current.prev
		}
		l.length--
		return
	}
}
