package main

import "fmt"

func main() {
	fmt.Println(generatePossibleNextMoves("++++++"))
}

func generatePossibleNextMoves(currentState string) []string {
	result := make([]string, 0)
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			tmp := make([]byte, len(currentState))
			copy(tmp, currentState)
			tmp[i] = '-'
			tmp[i+1] = '-'
			result = append(result, string(tmp))
		}
	}

	return result
}
