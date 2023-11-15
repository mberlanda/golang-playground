package exercises

func groupAnagrams(strs []string) [][]string {
	// Write a Go function 'groupAnagrams' that groups anagrams from a given slice of strings,
	// where anagrams are defined as words or phrases with the same letters, regardless of order.
	table := make(map[int][]string)

	for _, word := range strs {
		var count [26]int
		for _, w := range word {
			count[w-'a']++
		}
		hash := hashCount(count)
		table[hash] = append(table[hash], word)
	}

	result := make([][]string, 0, len(table))
	for _, group := range table {
		result = append(result, group)
	}
	return result
}

func hashCount(count [26]int) int {
	hash := 0
	for _, val := range count {
		hash = hash*31 + val
	}
	return hash
}
