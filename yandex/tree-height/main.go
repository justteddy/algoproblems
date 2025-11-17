package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	fields := strings.Fields(line)

	bst := Bst{}
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err == nil {
			bst.Add(num)
		}
	}

	fmt.Println(bst.height)
}

type Bst struct {
	root   *Node
	height int
}

type Node struct {
	left, right *Node
	val         int
}

func (t *Bst) Add(val int) {
	if t.root == nil {
		t.root = &Node{val: val}
		t.height = 0
		return
	}

	depth := 0
	var parent *Node
	curr := t.root

	for curr != nil {
		if val == curr.val {
			return
		}

		parent = curr
		if val < curr.val {
			curr = curr.left
		} else if val > curr.val {
			curr = curr.right
		}
		depth++
	}

	if val < parent.val {
		parent.left = &Node{val: val}
	} else {
		parent.right = &Node{val: val}
	}

	if depth > t.height {
		t.height = depth
	}
}
