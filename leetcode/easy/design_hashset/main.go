package main

// https://leetcode.com/problems/design-hashset/
func main() {

}

const bucketSize = 1e06 + 1

type MyHashSet struct {
	buckets [bucketSize]*int
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets: [bucketSize]*int{},
	}
}

func (h *MyHashSet) Add(key int) {
	if h.buckets[key] == nil {
		h.buckets[key] = &key
		return
	}
}

func (h *MyHashSet) Remove(key int) {
	h.buckets[key] = nil
}

func (h *MyHashSet) Contains(key int) bool {
	return h.buckets[key] != nil
}
