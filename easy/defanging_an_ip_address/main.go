package main

import (
	"fmt"
	"strings"
)

/*
	Given a valid (IPv4) IP address, return a defanged version of that IP address.

	A defanged IP address replaces every period "." with "[.]".
	Example 1:

	Input: address = "1.1.1.1"
	Output: "1[.]1[.]1[.]1"
	Example 2:

	Input: address = "255.100.50.0"
	Output: "255[.]100[.]50[.]0"

	Constraints:
	The given address is a valid IPv4 address.
*/

func main() {
	fmt.Println(defangIPaddr("255.100.50.0"))
}

// simplest solution
func defangIPaddrSimple(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

// performance solution
func defangIPaddr(address string) string {
	result := make([]byte, 0, len(address)+6)
	for i, j := 0, 0; i < len(address); i, j = i+1, j+1 {
		if address[i] == 46 { // it's byte code of period "."
			result = append(result, []byte{91, 46, 93}...)
			j += 2
			continue
		}
		result = append(result, address[i])
	}

	return string(result)
}
