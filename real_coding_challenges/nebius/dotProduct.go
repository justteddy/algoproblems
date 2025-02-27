package nebius

import "fmt"

// Даны 2 вектора целых чисел одинаковой длины, заданные в сжатой форме списками пар вида `(value, count)` .
// Например, вектор `[4, 4, 5]` задается как `[(4, 2), (5, 1)]` .
// Необходимо посчитать скалярное произведение заданных векторов.
// Пример:
// 1 1 1
// 1 1 10
// `multiply([(1, 3)], [(1, 2), (10, 1)]) -> 12

func main() {
	fmt.Println(multiply([][]int{{1, 3}}, [][]int{{1, 2}, {10, 1}}))
}

func multiply(left, right [][]int) int {
	sum := 0
	lIdx, rIdx := 0, 0
	lVal, rVal := left[lIdx][0], right[rIdx][0]
	lCnt, rCnt := left[lIdx][1], right[rIdx][1]

	for {
		// если равны кол-ва
		if lCnt == rCnt {
			sum += lVal * rVal * lCnt
			lIdx++
			rIdx++
			if lIdx >= len(left) || rIdx >= len(right) {
				break
			}

			lVal, rVal = left[lIdx][0], right[rIdx][0]
			lCnt, rCnt = left[lIdx][1], right[rIdx][1]
			continue
		}

		minCnt := min(lCnt, rCnt)
		sum += lVal * rVal * minCnt

		if lCnt > rCnt {
			lCnt = lCnt - rCnt

			rIdx++
			// проверить длину слайса - не нужна проверка
			rVal = right[rIdx][0]
			rCnt = right[rIdx][1]
		} else {
			rCnt = rCnt - lCnt

			lIdx++
			// проверить длину слайса - не нужна проверка
			lVal = left[lIdx][0]
			lCnt = left[lIdx][1]
		}
	}

	return sum
}
