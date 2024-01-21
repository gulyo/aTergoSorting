package main

import (
	"strconv"
	"strings"
)

type WordWeight struct {
	word    Word
	symbols []string
	weights []float32
	missing []string
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
	return strconv.Itoa(actual.word.index) + " " +
		string(actual.word.text) + "," +
		strings.Join(actual.missing, "|")
}
