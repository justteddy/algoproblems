package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	m := make(map[string]bool)
	for sc.Scan() {
		m[sc.Text()] = true
	}
	fmt.Println(len(m))
}
