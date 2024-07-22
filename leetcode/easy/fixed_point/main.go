package main

func fixedPoint(arr []int) int {
	res := -1
	mid, lo, hi := 0, 0, len(arr)-1
	for lo <= hi {
		mid = (lo + hi) / 2
		if arr[mid] == mid {
			res = mid
			// check only left side then
			hi = mid - 1
		} else {
			if arr[mid] > mid {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}

	return res
}
