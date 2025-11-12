package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val  int
}

func main() {
	fmt.Println(prefixSum([]int{0, 1, 2, 3, 4, 5, 6}))
	fmt.Println(twoPointers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 13))
	fmt.Println(slidingWindow([]int{2, 1, 5, 1, 3, 2}, 3))

	list := Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val: 5,
					},
				},
			},
		},
	}

	// cycle
	list.Next.Next.Next = list.Next.Next

	fmt.Println(fastAndSlowPointers(&list))
}

func prefixSum(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}

	return res
}

func twoPointers(sorted []int, target int) (int, int) {
	for i, j := 0, len(sorted)-1; i < j; {
		sum := sorted[i] + sorted[j]
		if sum > target {
			j--
			continue
		}

		if sum < target {
			i++
			continue
		}

		return i, j
	}

	return -1, -1
}

func slidingWindow(nums []int, k int) int {
	currentSum := 0
	maxSum := 0
	for i, v := range nums {
		if i < k {
			currentSum += v
			maxSum = currentSum
			continue
		}

		currentSum -= nums[i-k]
		currentSum += v
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func fastAndSlowPointers(list *Node) bool {
	fast, slow := list, list
	for {
		if fast == nil || fast.Next == nil {
			break
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
