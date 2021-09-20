package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/number-of-recent-calls/
func main() {
	rc := Constructor()
	fmt.Println(
		rc.Ping(1),
		rc.Ping(100),
		rc.Ping(3001),
		rc.Ping(3002),
	)
}

type RecentCounter struct {
	window *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{
		window: list.New(),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.window.PushBack(t)
	for rc.window.Front().Value.(int) < t-3000 {
		rc.window.Remove(rc.window.Front())
	}
	return rc.window.Len()
}
