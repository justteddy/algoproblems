package main

import "fmt"

/*
	Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

	Implement the MovingAverage class:

	MovingAverage(int size) Initializes the object with the size of the window size.
	double next(int val) Returns the moving average of the last size values of the stream.


	Example 1:

	Input
	["MovingAverage", "next", "next", "next", "next"]
	[[3], [1], [10], [3], [5]]
	Output
	[null, 1.0, 5.5, 4.66667, 6.0]

	Explanation
	MovingAverage movingAverage = new MovingAverage(3);
	movingAverage.next(1); // return 1.0 = 1 / 1
	movingAverage.next(10); // return 5.5 = (1 + 10) / 2
	movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
	movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3
*/

type MovingAverage struct {
	size      int
	windowSum float64
	els       []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:      size,
		windowSum: 0,
		els:       make([]int, 0, size),
	}
}

func (m *MovingAverage) Next(val int) float64 {
	if len(m.els) == m.size {
		m.windowSum -= float64(m.els[0])
		m.els = m.els[1:]
	}

	m.els = append(m.els, val)
	m.windowSum += float64(val)

	return m.windowSum / float64(len(m.els))

}

func main() {
	m := Constructor(3)
	fmt.Println(m.Next(1))
	fmt.Println(m.Next(10))
	fmt.Println(m.Next(3))
	fmt.Println(m.Next(5))
}
