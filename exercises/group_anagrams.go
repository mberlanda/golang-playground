package exercises

import (
	"runtime"
	"sync"
)

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

// With small inputs the following approach leads to parallelism overhead
func groupAnagramsGoRoutines(strs []string) [][]string {
	var wg sync.WaitGroup
	wg.Add(len(strs))

	table := make(map[int][]string)
	var mutex sync.Mutex

	for _, s := range strs {
		go func(word string) {
			defer wg.Done()
			var count [26]int
			for _, w := range word {
				count[w-'a']++
			}
			hash := hashCount(count)
			mutex.Lock()
			table[hash] = append(table[hash], word)
			mutex.Unlock()
		}(s)
	}

	wg.Wait()

	result := make([][]string, 0, len(table))
	for _, group := range table {
		result = append(result, group)
	}
	return result
}

func groupAnagramsPool(strs []string) [][]string {
	var wg sync.WaitGroup
	wg.Add(len(strs))

	table := make(map[int][]string)
	var mutex sync.Mutex

	numWorkers := runtime.NumCPU()
	pool := make(chan struct{}, numWorkers)
	for i := 0; i < numWorkers; i++ {
		pool <- struct{}{}
	}
	for _, s := range strs {
		<-pool // Acquire a worker from the pool
		go func(word string) {
			defer wg.Done()
			var count [26]int
			for _, w := range word {
				count[w-'a']++
			}
			hash := hashCount(count)
			mutex.Lock()
			table[hash] = append(table[hash], word)
			mutex.Unlock()
			pool <- struct{}{} // Release the worker back to the pool
		}(s)
	}

	wg.Wait()

	result := make([][]string, 0, len(table))
	for _, group := range table {
		result = append(result, group)
	}
	return result
}
