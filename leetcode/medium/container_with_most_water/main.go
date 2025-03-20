package main

func main() {

}

// https://leetcode.com/problems/container-with-most-water
func maxArea(height []int) int {
	maxVol := 0
	left, right := 0, len(height)-1

	for left != right {
		maxVol = max(maxVol, min(height[left], height[right])*(right-left))

		if height[left] <= height[right] {
			left++
			continue
		}
		right--
	}
	return maxVol
}

// naive solution
func maxArea2(height []int) int {
	maxVol := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			maxVol = max(maxVol, min(height[i], height[j])*(j-i))
		}
	}

	return maxVol
}
