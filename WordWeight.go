package main

import (
	"fmt"
	"strconv"
	"strings"
)

var PrintWidthWord int = 8
var PrintWidthRowNum int = 8

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

func (actual WordWeight) String(rPad bool) string {
	var weightPrint []string = make([]string, len(actual.weights))
	for i := range actual.weights {
		weightPrint[i] = fmt.Sprintf("%.0f", actual.weights[i])
	}
	if rPad {
		return fmt.Sprintf("%"+strconv.Itoa(PrintWidthRowNum)+"d", actual.word.index) + "," +
			fmt.Sprintf("%"+strconv.Itoa(PrintWidthWord)+"s", string(actual.word.text)) + "," +
			strings.Join(weightPrint, ":") + "," +
			strings.Join(actual.missing, "|")
	} else {
		return fmt.Sprintf(strconv.Itoa(actual.word.index)) + "," +
			fmt.Sprintf(string(actual.word.text)) + "," +
			strings.Join(weightPrint, ":") + "," +
			strings.Join(actual.missing, "|")
	}

}
