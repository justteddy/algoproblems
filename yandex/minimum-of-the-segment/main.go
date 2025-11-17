package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	if _, err := fmt.Fscan(in, &n, &k); err != nil {
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// deque будет хранить индексы элементов массива a
	deque := make([]int, 0, n)

	for i := 0; i < n; i++ {
		// 1. Удаляем из головы элементы, которые вышли за левую границу окна
		// Текущее окно: [i-k+1, i], значит индексы <= i-k нам не подходят
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 2. Удаляем из хвоста все элементы, которые >= текущего a[i]
		// чтобы поддерживать монотонно возрастающую очередь по значениям
		for len(deque) > 0 && a[deque[len(deque)-1]] >= a[i] {
			deque = deque[:len(deque)-1]
		}

		// 3. Добавляем текущий индекс
		deque = append(deque, i)

		// 4. Как только окно стало "полным" (i >= k-1), печатаем минимум
		if i >= k-1 {
			minIndex := deque[0]
			fmt.Fprintln(out, a[minIndex])
		}
	}
}
