package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(reader, &N)
	dict := make(map[string]string, N*2)

	for range N {
		var w1, w2 string
		fmt.Fscan(reader, &w1)
		fmt.Fscan(reader, &w2)
		dict[w1] = w2
		dict[w2] = w1
	}

	var s string
	fmt.Fscan(reader, &s)

	fmt.Println(dict[s])
}
