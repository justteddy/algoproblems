/*

You are in charge of a display advertising program. Your ads are displayed on websites all over the internet. You have some CSV input data that counts how many times that users have clicked on an ad on each individual domain. Every line consists of a click count and a domain name, like this:

counts = [ "900,google.com",
     "60,mail.yahoo.com",
     "10,mobile.sports.yahoo.com",
     "40,sports.yahoo.com",
     "300,yahoo.com",
     "10,stackoverflow.com",
     "20,overflow.com",
     "5,com.com",
     "2,en.wikipedia.org",
     "1,m.wikipedia.org",
     "1,mobile.sports",
     "1,google.co.uk"]

Write a function that takes this input as a parameter and returns a data structure containing the number of clicks that were recorded on each domain AND each subdomain under it. For example, a click on "mail.yahoo.com" counts toward the totals for "mail.yahoo.com", "yahoo.com", and "com". (Subdomains are added to the left of their parent domain. So "mail" and "mail.yahoo" are not valid domains. Note that "mobile.sports" appears as a separate domain near the bottom of the input.)

Sample output (in any order/format):

calculateClicksByDomain(counts) =>
com:                     1345
google.com:              900
stackoverflow.com:       10
overflow.com:            20
yahoo.com:               410
mail.yahoo.com:          60
mobile.sports.yahoo.com: 10
sports.yahoo.com:        50
com.com:                 5
org:                     3
wikipedia.org:           3
en.wikipedia.org:        2
m.wikipedia.org:         1
mobile.sports:           1
sports:                  1
uk:                      1
co.uk:                   1
google.co.uk:            1

n: number of domains in the input
(individual domains and subdomains have a constant upper length)


*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	counts := []string{
		"900,google.com",
		"60,mail.yahoo.com",
		"10,mobile.sports.yahoo.com",
		"40,sports.yahoo.com",
		"300,yahoo.com",
		"10,stackoverflow.com",
		"20,overflow.com",
		"5,com.com",
		"2,en.wikipedia.org",
		"1,m.wikipedia.org",
		"1,mobile.sports",
		"1,google.co.uk",
	}

	calculateClicksByDomain(counts)
}

// simple solution which I could use... LOL
func calculateClicksByDomain(counts []string) {
	hm := make(map[string]int, 0)
	for _, siteCount := range counts {
		countAndSite := strings.Split(siteCount, ",")
		count, _ := strconv.Atoi(countAndSite[0])

		domains := strings.Split(countAndSite[1], ".")

		for i := len(domains) - 1; i >= 0; i-- {
			str := strings.Join(domains[i:], ".")
			if _, ok := hm[str]; !ok {
				hm[str] = count
				continue
			}
			hm[str] += count
		}
	}

	for d, count := range hm {
		fmt.Printf("%s: %d \n", d, count)
	}
}

// trie solution which is overhead for that task ... :)
func calculateClicksByDomain_(counts []string) {
	trie := NewTrie()
	for _, siteCount := range counts {
		parts := strings.Split(siteCount, ",")
		count, _ := strconv.Atoi(parts[0])
		trie.Insert(parts[1], count)
	}

	trie.Traverse()
}

type Node struct {
	Children map[string]*Node
	Domain   string
	Count    int
}

func NewTrie() *Node {
	return &Node{
		Children: make(map[string]*Node, 0),
	}
}

func (n *Node) Insert(domain string, count int) {
	parts := strings.Split(domain, ".")

	curr := n
	for i := len(parts) - 1; i >= 0; i-- {
		str := strings.Join(parts[i:], ".")
		child, ok := curr.Children[str]
		if !ok {
			child := &Node{
				Domain:   str,
				Count:    count,
				Children: make(map[string]*Node),
			}
			curr.Children[str] = child
			curr = child
			continue
		}
		child.Count += count
		curr = child
	}
}

func (n *Node) Traverse() {
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if n == nil {
			return
		}

		fmt.Printf("%s: %d \n", node.Domain, node.Count)
		for _, child := range node.Children {
			traverse(child)
		}
	}

	traverse(n)
}
