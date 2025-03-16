package main

import "fmt"

// https://leetcode.com/problems/asteroid-collision
func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
}

func asteroidCollision(asteroids []int) []int {
	var st stack
	st.Push(asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		aster := asteroids[i]
		if aster > 0 {
			st.Push(aster)
			continue
		}
		if aster < 0 {
			for {
				if st.Len() == 0 || st.Peek() < 0 {
					st.Push(aster)
					break
				}
				diff := aster + st.Peek() // peek is positive here
				if diff == 0 {
					st.Pop()
					break
				} else if diff < 0 {
					st.Pop()
					continue
				}
				break
			}
		}
	}

	return st
}

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last, true
}

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) Len() int { return len(*s) }
