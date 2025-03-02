package crawler_log_folder

// https://leetcode.com/problems/crawler-log-folder/
func minOperations(logs []string) int {
	curr := 0
	for _, op := range logs {
		switch op {
		case "../":
			if curr != 0 {
				curr--
			}
		case "./":
			continue
		default:
			curr++
		}
	}

	return curr
}
