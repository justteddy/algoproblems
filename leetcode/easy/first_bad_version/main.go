package main

https://leetcode.com/problems/first-bad-version/
func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool { return true }

func firstBadVersion(n int) int {
	lo, hi := 1, n
	mid := 0
	for lo < hi {
		mid = (lo + hi) / 2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
