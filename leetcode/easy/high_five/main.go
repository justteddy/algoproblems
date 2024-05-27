package main

import (
	"fmt"
)

func main() {
	fmt.Println(highFive([][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}}))
}

func highFive(items [][]int) [][]int {
	sorted := make(map[int][101]int)
	for _, studScore := range items {
		studentID, score := studScore[0], studScore[1]
		scores, ok := sorted[studentID]
		if !ok {
			tmp := [101]int{}
			tmp[score] = 1
			sorted[studentID] = tmp
			continue
		}
		scores[score]++
		sorted[studentID] = scores
	}

	res := make([][]int, 0, len(sorted))
	for studID := 1; studID <= 1000; studID++ {
		if len(res) == len(sorted) {
			break
		}

		if _, ok := sorted[studID]; !ok {
			continue
		}

		scores := sorted[studID]
		sum, cnt := 0, 5
		for j := 100; cnt > 0; j-- {
			if scores[j] == 0 {
				continue
			}

			if scores[j] > cnt {
				sum += j * cnt
			} else {
				sum += j * scores[j]
			}
			cnt -= scores[j]
		}
		res = append(res, []int{studID, sum / 5})
	}

	return res
}
