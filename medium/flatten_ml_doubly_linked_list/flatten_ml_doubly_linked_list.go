package flatten_ml_doubly_linked_list

/*
   You are given a doubly linked list which in addition to the next and previous pointers, it could have a child pointer, which may or may not point to a separate doubly linked list. These child lists may have one or more children of their own, and so on, to produce a multilevel data structure, as shown in the example below.
   Flatten the list so that all the nodes appear in a single-level, doubly linked list. You are given the head of the first level of the list.

   Example 1:
   Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
   Output: [1,2,3,7,8,11,12,9,10,4,5,6]

    Example 2:
    Input: head = [1,2,null,3]
    Output: [1,3,2]
    Explanation:
    The input multilevel linked list is as follows:

      1---2---NULL
      |
      3---NULL

    Example 3:
     1---2---3---4---5---6--NULL
             |
             7---8---9---10--NULL
                 |
                 11--12--NULL

    Output: [1,2,3,7,8,11,12,9,10,4,5,6]

*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	head, _ := mlTraverse(root, nil)
	return head
}

func mlTraverse(head *Node, prev *Node) (*Node, *Node) {
	cur := head
	for cur != nil {
		cur.Prev = prev
		if cur.Child != nil {
			nestedHead, nestedLast := mlTraverse(cur.Child, nil)
			cur.Child = nil

			nestedHead.Prev = cur
			nestedLast.Next = cur.Next
			if cur.Next != nil {
				cur.Next.Prev = nestedLast
			}

			tmpNext := cur.Next
			cur.Next = nestedHead
			cur = tmpNext
			prev = nestedLast
			continue
		}
		prev = cur
		cur = cur.Next
	}

	return head, prev
}
