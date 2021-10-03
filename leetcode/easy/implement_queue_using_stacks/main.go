package main

// https://leetcode.com/problems/implement-queue-using-stacks/
func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
}

type MyQueue struct {
	push *Stack
	pull *Stack
	len  int
}

func Constructor() MyQueue {
	return MyQueue{
		push: new(Stack),
		pull: new(Stack),
	}
}

func (q *MyQueue) Push(x int) {
	defer func() { q.len++ }()
	if q.pull.IsEmpty() {
		q.push.Push(x)
	} else {
		q.resetPush()
		for !q.pull.IsEmpty() {
			q.push.Push(q.pull.Pop())
		}
		q.push.Push(x)
		q.resetPull()
	}
}

func (q *MyQueue) Pop() int {
	defer func() { q.len-- }()
	if q.pull.IsEmpty() {
		for !q.push.IsEmpty() {
			q.pull.Push(q.push.Pop())
		}
		return q.pull.Pop()
	} else {
		return q.pull.Pop()
	}
}

func (q *MyQueue) Peek() int {
	if q.pull.IsEmpty() {
		for !q.push.IsEmpty() {
			q.pull.Push(q.push.Pop())
		}
		return (*q.pull)[q.pull.Len()-1]
	} else {
		return (*q.pull)[q.pull.Len()-1]
	}
}

func (q *MyQueue) Empty() bool {
	return q.len == 0
}

func (q *MyQueue) resetPush() {
	q.push = new(Stack)
}

func (q *MyQueue) resetPull() {
	q.pull = new(Stack)
}

type Stack []int

func (s *Stack) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack) Len() int      { return len(*s) }
func (s *Stack) Push(el int)   { *s = append(*s, el) }
func (s *Stack) Pop() int {
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}
