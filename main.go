package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var numbers []int64
	for scanner.Scan() {
		var num int64
		fmt.Sscan(scanner.Text(), &num)
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if len(numbers) < 2 {
		fmt.Println("List must contain at least two numbers.")
		return
	}

	maxP1, maxP2 := math.MinInt64, math.MinInt64
	maxN1, maxN2 := math.MaxInt64, math.MaxInt64
	for _, n := range numbers {
		n := int(n)
		if n > maxP1 {
			maxP2 = maxP1
			maxP1 = n
		} else if n > maxP2 {
			maxP2 = n
		}

		if n < maxN1 {
			maxN2 = maxN1
			maxN1 = n
		} else if n < maxN2 {
			maxN2 = n
		}
	}

	pMax := maxP1 * maxP2
	nMax := maxN1 * maxN2

	if pMax > nMax {
		fmt.Println(maxP2, maxP1)
	} else {
		fmt.Println(maxN1, maxN2)
	}
}
