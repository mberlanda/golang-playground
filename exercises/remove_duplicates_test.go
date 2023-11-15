package exercises

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
		message  string
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5, "Duplicates in the middle of the array"},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, "Duplicates at the beginning and end"},
		{[]int{1, 2, 3}, 3, "No duplicates"},
		{[]int{4, 4, 4, 4}, 2, "All elements are same"},
	}

	for _, tc := range testCases {
		t.Run(tc.message, func(t *testing.T) {
			result := removeDuplicates(tc.nums)
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %d, got %d", tc.message, tc.expected, result)
			}
		})
	}
}
