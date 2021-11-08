package main

import (
	"container/list"
	"fmt"
	"hash/fnv"
)

// https://leetcode.com/problems/design-hashmap/
func main() {

}

const bucketSize = 1024

type KeyValue struct {
	Key   int
	Value int
}

type MyHashMap struct {
	buckets [bucketSize]*list.List
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: [1024]*list.List{},
	}
}

func (h *MyHashMap) Put(key int, value int) {
	bucket := hashFn(key) % bucketSize
	if h.buckets[bucket] == nil {
		h.buckets[bucket] = list.New()
		h.buckets[bucket].PushFront(KeyValue{
			Key:   key,
			Value: value,
		})
	} else {
		curr := h.buckets[bucket].Front()
		for curr != nil {
			if curr.Value.(KeyValue).Key == key {
				curr.Value = KeyValue{
					Key:   key,
					Value: value,
				}
				return
			}
			curr = curr.Next()
		}
		h.buckets[bucket].PushFront(KeyValue{
			Key:   key,
			Value: value,
		})
	}
}

func (h *MyHashMap) Get(key int) int {
	bucket := hashFn(key) % bucketSize
	if h.buckets[bucket] == nil {
		return -1
	} else {
		curr := h.buckets[bucket].Front()
		for curr != nil {
			if curr.Value.(KeyValue).Key == key {
				return curr.Value.(KeyValue).Value
			}
			curr = curr.Next()
		}
	}
	return -1
}

func (h *MyHashMap) Remove(key int) {
	bucket := hashFn(key) % bucketSize
	if h.buckets[bucket] == nil {
		return
	}
	curr := h.buckets[bucket].Front()
	for curr != nil {
		if curr.Value.(KeyValue).Key == key {
			h.buckets[bucket].Remove(curr)
		}
		curr = curr.Next()
	}
}

func hashFn(v int) uint32 {
	hasher := fnv.New32a()
	hasher.Write([]byte(fmt.Sprintf("%d", v)))
	return hasher.Sum32()
}
