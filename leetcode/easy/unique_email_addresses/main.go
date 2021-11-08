package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/unique-email-addresses/
func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}))
}

func numUniqueEmails(emails []string) int {
	unique := make(map[string]struct{})
	for _, email := range emails {
		unique[fmtEmail(email)] = struct{}{}
	}

	return len(unique)
}

func fmtEmail(email string) string {
	parts := strings.Split(email, "@")

	b := new(strings.Builder)
	for i := 0; i < len(parts[0]); i++ {
		if parts[0][i] == '.' {
			continue
		}
		if parts[0][i] == '+' {
			break
		}
		b.WriteByte(parts[0][i])
	}

	return b.String() + "@" + parts[1]
}
