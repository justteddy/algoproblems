package main

import "fmt"

func main() {
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "practice"))
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	res := len(wordsDict)
	w1Idx, w2Idx := -1, -1
	for i, w := range wordsDict {
		if w == word1 {
			w1Idx = i
		}
		if w == word2 {
			w2Idx = i
		}

		if w1Idx != -1 && w2Idx != -1 {
			dist := abs(w1Idx - w2Idx)
			if dist < res {
				res = dist
			}
		}
	}

	return res
}

func shortestDistance_(wordsDict []string, word1 string, word2 string) int {
	w1Idxs, w2Idxs := make([]int, 0), make([]int, 0)
	for i, w := range wordsDict {
		if w == word1 {
			w1Idxs = append(w1Idxs, i)
		}
		if w == word2 {
			w2Idxs = append(w2Idxs, i)
		}
	}

	res := len(wordsDict)
	for _, w1Idx := range w1Idxs {
		for _, w2Idx := range w2Idxs {
			dist := abs(w1Idx - w2Idx)
			if dist < res {
				res = dist
			}
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
