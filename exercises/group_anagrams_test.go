package exercises

import (
	"reflect"
	"sort"
	"testing"
)

func sortSliceOfSliceOfStrings(input [][]string) [][]string {
	customSort := func(slice []string) {
		sort.Strings(slice)
	}
	for _, group := range input {
		customSort(group)
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i][0] < input[j][0]
	})
	return input
}

func are2DSlicesEqual(a, b [][]string) bool {
	sortedA := sortSliceOfSliceOfStrings(a)
	sortedB := sortSliceOfSliceOfStrings(b)
	return reflect.DeepEqual(sortedA, sortedB)
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		// Add more test cases here if needed
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := groupAnagrams(test.strs)
			if !are2DSlicesEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
