package main

import "fmt"

// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}

// monotonic stack
func finalPrices(prices []int) []int {
	stack := new(Stack)
	for i := len(prices) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && stack.Peek() > prices[i] {
			stack.Pop()
		}

		if stack.IsEmpty() {
			stack.Push(prices[i])
		} else {
			tmp := prices[i]
			prices[i] = prices[i] - stack.Peek()
			stack.Push(tmp)
		}
	}

	return prices
}

type Stack []int

func (s *Stack) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack) Len() int      { return len(*s) }
func (s *Stack) Peek() int     { return (*s)[len(*s)-1] }
func (s *Stack) Push(el int)   { *s = append(*s, el) }
func (s *Stack) Pop() int {
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}

// naive solution
func finalPrices_(prices []int) []int {
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] = prices[i] - prices[j]
				break
			}
		}
	}
	return prices
}
