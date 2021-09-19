package main

import "fmt"

// https://leetcode.com/problems/destination-city/
func main() {
	fmt.Println(destCity([][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}))
}

func destCity(paths [][]string) string {
	cities := make(map[string]string, len(paths))
	for _, path := range paths {
		cities[path[0]] = path[1]
	}

	for _, dest := range cities {
		if _, ok := cities[dest]; !ok {
			return dest
		}
	}

	return ""
}

func destCity_(paths [][]string) string {
	hm := make(map[string]int)
	for _, path := range paths {
		if _, ok := hm[path[0]]; !ok {
			hm[path[0]] = 1
		} else {
			delete(hm, path[0])
		}

		if _, ok := hm[path[1]]; !ok {
			hm[path[1]] = 2
		}
	}

	for key, val := range hm {
		if val == 2 {
			return key
		}
	}

	return ""
}
