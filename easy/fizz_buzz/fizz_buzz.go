package fizz_buzz

import (
	"strconv"
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
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		mod3 := i % 3
		mod5 := i % 5
		if mod3 == 0 && mod5 == 0 {
			res = append(res, "FizzBuzz")
			continue
		}
		if mod3 == 0 {
			res = append(res, "Fizz")
			continue
		}
		if mod5 == 0 {
			res = append(res, "Buzz")
			continue
		}
		res = append(res, strconv.Itoa(i))
	}
	return res
}
