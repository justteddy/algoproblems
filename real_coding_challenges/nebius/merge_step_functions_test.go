package nebius_test

import (
	"fmt"
	"testing"
	"sort"
)

type Point struct {
	Time  int
	Value int
}

/*
	У нас есть две ступенчатые функции, каждая представлена списком точек {Time, Value}.
	Значение функции не меняется между точками измерения (стандартный “степ-вью” график).
	Нужно построить новую ступенчатую функцию, которая будет суммой двух исходных.
*/
func MergeStepFunctions(func1, func2 []Point) []Point {
	// Создаем набор всех уникальных временных точек
	timeMap := make(map[int]bool)
	for _, p := range func1 {
		timeMap[p.Time] = true
	}
	for _, p := range func2 {
		timeMap[p.Time] = true
	}

	// Преобразуем в слайс и сортируем
	times := make([]int, 0, len(timeMap))
	for t := range timeMap {
		times = append(times, t)
	}
	sort.Ints(times)

	// Для каждой временной точки находим значения обеих функций и суммируем
	result := make([]Point, 0, len(times))
	var value1, value2 int
	var idx1, idx2 int

	for _, t := range times {
		// Обновляем значение первой функции, если достигли новой точки измерения
		for idx1 < len(func1) && func1[idx1].Time <= t {
			value1 = func1[idx1].Value
			idx1++
		}

		// Обновляем значение второй функции, если достигли новой точки измерения
		for idx2 < len(func2) && func2[idx2].Time <= t {
			value2 = func2[idx2].Value
			idx2++
		}

		// Добавляем точку с суммой значений в результат
		result = append(result, Point{Time: t, Value: value1 + value2})
	}

	return result
}

func Test_MergeStepFunctions(t *testing.T) {
	f1 := []Point{{1, 3}, {4, 5}, {7, 2}}
	f2 := []Point{{2, 4}, {6, 1}}

	result := MergeStepFunctions(f1, f2)
	for _, p := range result {
		fmt.Printf("Time: %d, Value: %d\n", p.Time, p.Value)
	}
}
