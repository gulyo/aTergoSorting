package main

type WordWeight struct {
	word    Word
	weights []float32
}

func (actual WordWeight) lessThan(wordB WordWeight) bool {
	a := actual.weights
	b := wordB.weights
	endA := len(a) - 1
	endB := len(b) - 1
	length := min(endA, endB) + 1

	var i int = 0
	for i < length {
		if a[endA-i] > b[endB-i] {
			return false
		}
		if a[endA-i] < b[endB-i] {
			return true
		}
		i++
	}
	if endA > endB {
		return false
	}
	if endA < endB {
		return true
	}
	return false
}

func (actual WordWeight) String() string {
	return string(actual.word)
}
