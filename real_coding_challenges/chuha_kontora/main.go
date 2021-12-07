package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(smallestNegativeBalance([][]string{
		{"Alex", "Blake", "2"},
		{"Blake", "Alex", "2"},
		{"Casey", "Alex", "5"},
		{"Blake", "Casey", "7"},
		{"Alex", "Blake", "4"},
		{"Alex", "Casey", "4"},
	}))
}

func smallestNegativeBalance(debts [][]string) []string {
	debtsMap := make(map[string]int)
	for _, debt := range debts {
		amount, _ := strconv.Atoi(debt[2])
		debtsMap[debt[0]] -= amount
		debtsMap[debt[1]] += amount
	}

	var result []string
	var smallest = -1
	for name, amount := range debtsMap {
		if amount < smallest {
			result = []string{name}
			smallest = amount
		} else if amount == smallest {
			result = append(result, name)
		}
	}

	if len(result) == 0 {
		return []string{"No one has a negative balance"}
	}

	sort.Strings(result)

	return result
}
