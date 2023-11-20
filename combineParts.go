package main

import (
	"sync"
)

func combineParts(c *[]*[]WordWeight) *[]WordWeight {
	var container *[]*[]WordWeight = c
	var wg sync.WaitGroup
	var wgCollect sync.WaitGroup
	for len(*container) > 1 {
		var tmpContainer []*[]WordWeight
		var channel chan *[]WordWeight = make(chan *[]WordWeight)

		wgCollect.Add(1)
		go func() {
			defer wgCollect.Done()
			for combedResult := range channel {
				tmpContainer = append(tmpContainer, combedResult)
			}
		}()

		var i int = 0
		for i < (len(*container) - 1) {
			wg.Add(1)
			go func(a, b *[]WordWeight) {
				defer wg.Done()
				channel <- combSorted(a, b)
			}((*container)[i], (*container)[i+1])
			i = i + 2
		}
		wg.Wait()
		close(channel)
		wgCollect.Wait()
		if i < len(*container) {
			tmpContainer = append(tmpContainer, (*container)[i])
		}
		container = &tmpContainer
	}
	return (*container)[0]
}
