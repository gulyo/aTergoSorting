package main

import (
	"os"
	"sort"
	"sync"
)

func sortedParts() *[]*[]WordWeight {
	var container []*[]WordWeight
	abc := readAbc(os.Args[1])
	var wg sync.WaitGroup
	for w := range readNames(os.Args[2]) {
		wg.Add(1)
		go func(words *[]Word) {
			defer func() { wg.Done() }()

			var result []WordWeight
			for _, word := range *words {
				if word != "" {
					result = append(result, word.CalculateWeights(&abc))
				}
			}
			sort.Slice(result, lessComparator(result))
			container = append(container, &result)
		}(w)
	}
	wg.Wait()
	return &container
}
