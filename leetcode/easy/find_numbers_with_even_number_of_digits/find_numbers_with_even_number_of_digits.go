package find_numbers_with_even_number_of_digits

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	result := 0
	for _, n := range nums {
		if numOfDigits(n)&1 == 0 { // even
			result++
		}
	}

	return result
}

func numOfDigits(n int) (result int) {
	for n != 0 {
		n /= 10
		result++
	}
	return
}
