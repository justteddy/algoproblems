package three_consecutive_odds

// https://leetcode.com/problems/three-consecutive-odds/
func threeConsecutiveOdds(arr []int) bool {
	cnt := 0
	for _, v := range arr {
		if v%2 == 0 {
			cnt = 0
			continue
		}

		cnt++
		if cnt == 3 {
			return true
		}
	}

	return false
}
