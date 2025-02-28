package fair_candy_swap

// https://leetcode.com/problems/fair-candy-swap/
// naive solution
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum, bobSum := 0, 0
	for _, v := range aliceSizes {
		aliceSum += v
	}
	for _, v := range bobSizes {
		bobSum += v
	}

	for _, as := range aliceSizes {
		for _, bs := range bobSizes {
			optA := aliceSum - as + bs
			optB := bobSum - bs + as
			if optA == optB {
				return []int{as, bs}
			}
		}
	}

	return nil
}

func fairCandySwap2(aliceSizes []int, bobSizes []int) []int {
	aliceSum, bobSum := 0, 0
	for _, v := range aliceSizes {
		aliceSum += v
	}

	m := make(map[int]bool, len(bobSizes))
	for _, v := range bobSizes {
		bobSum += v
		m[v] = true
	}

	delta := (bobSum - aliceSum) / 2
	for _, v := range aliceSizes {
		if _, ok := m[v+delta]; ok {
			return []int{v, v + delta}
		}
	}

	return nil
}
