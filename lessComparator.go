package main

func lessComparator(words []WordWeight) func(wordI, wordJ int) bool {
	return func(wordI, wordJ int) bool {
		return words[wordI].lessThan(words[wordJ])
	}
}
