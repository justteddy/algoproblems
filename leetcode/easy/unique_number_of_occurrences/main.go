package main

// https://leetcode.com/problems/unique-number-of-occurrences
func main() {
	// Example usage
	arr := []int{1, 2, 2, 1, 1, 3}
	result := uniqueOccurrences(arr)
	println(result) // Output: true
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	k := make(map[int]bool)
	for _, cnt := range m {
		_, ok := k[cnt]
		if ok {
			return false
		}
		k[cnt] = true
	}

	return true
}
