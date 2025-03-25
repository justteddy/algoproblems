package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Operation struct {
	T int
	X int
	C rune
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// reading the length of the string
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// reading the string
	sStr, _ := reader.ReadString('\n')
	s := []rune(strings.TrimSpace(sStr))

	// reading the number of operations
	qStr, _ := reader.ReadString('\n')
	q, _ := strconv.Atoi(strings.TrimSpace(qStr))

	// reading the operations
	ops := make([]Operation, q)
	for i := 0; i < q; i++ {
		opStr, _ := reader.ReadString('\n')
		parts := strings.Split(opStr, " ")
		t, _ := strconv.Atoi(parts[0])
		x, _ := strconv.Atoi(parts[1])
		ops[i] = Operation{
			T: t,
			X: x,
			C: rune(parts[2][0]), // parts[2] is a single character 1 byte rune
		}

	}

	updateTime := make([]int, n)
	lastGlobalOpTime := 0
	globalOpType := 0

	// applying the operations
	for i, op := range ops {
		switch op.T {
		case 1:
			// replace character at position x with character c
			s[op.X-1] = op.C
			updateTime[op.X-1] = i
		case 2, 3:
			globalOpType = op.T
			lastGlobalOpTime = i
		}
	}

	for i := 0; i < n; i++ {
		if lastGlobalOpTime > updateTime[i] {
			if globalOpType == 2 {
				s[i] = unicode.ToLower(s[i])
			} else if globalOpType == 3 {
				s[i] = unicode.ToUpper(s[i])
			}
		}
	}

	fmt.Println(string(s))
}
