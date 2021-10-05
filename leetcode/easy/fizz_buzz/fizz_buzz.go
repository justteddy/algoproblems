package fizz_buzz

import (
	"fmt"
)

/*
	Write a program that outputs the string representation of numbers from 1 to n.
	But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

	Example:
	n = 15,

	Return:
	[
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz"
	]
*/

func fizzBuzz(n int) []string {
	answer := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		ans := ""
		if i%3 == 0 {
			ans += "Fizz"
		}
		if i%5 == 0 {
			ans += "Buzz"
		}
		if ans == "" {
			ans = fmt.Sprintf("%d", i)
		}
		answer = append(answer, ans)
	}
	return answer
}
