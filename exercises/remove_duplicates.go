package exercises

func removeDuplicates(nums []int) int {
	// Implement the logic to remove duplicates such that each element appears at most twice
	n := len(nums)
	if n <= 2 {
		return n
	}

	next := 2
	for curr := 2; curr < n; curr++ {
		if nums[curr] != nums[next-2] {
			nums[next] = nums[curr]
			next++
		}
	}
	return next
}
