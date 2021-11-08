package main

import (
	"container/list"
)

func main() {

}

const bucketSize = 1e05

type MyHashSet struct {
	buckets [bucketSize]*list.List
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets: [bucketSize]*list.List{},
	}
}

func (h *MyHashSet) Add(key int) {
	bucket := key % bucketSize
	if h.buckets[bucket] == nil {
		h.buckets[bucket] = list.New()
		h.buckets[bucket].PushFront(key)
		return
	}

	curr := h.buckets[bucket].Front()
	for curr != nil {
		if curr.Value.(int) == key {
			return
		}
	}
	h.buckets[bucket].PushFront(key)
}

func (h *MyHashSet) Remove(key int) {
	bucket := key % bucketSize
	if h.buckets[bucket] == nil {
		return
	}

	curr := h.buckets[bucket].Front()
	for curr != nil {
		if curr.Value.(int) == key {
			h.buckets[bucket].Remove(curr)
			return
		}
	}
}

func (h *MyHashSet) Contains(key int) bool {
	bucket := key % bucketSize
	if h.buckets[bucket] == nil {
		return false
	}

	curr := h.buckets[bucket].Front()
	for curr != nil {
		if curr.Value.(int) == key {
			return true
		}
	}

	return false
}
