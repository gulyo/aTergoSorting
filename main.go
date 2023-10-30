package main

import (
	"fmt"
	"os"
)

func main() {

	abc := readAbc(os.Args[1])

	for words := range readNames(os.Args[2]) {
		for _, word := range words {
			result := word.CalculateWeights(&abc)

			fmt.Println(*result.word, " -> ", *result.weights)
		}
		fmt.Println(*words)
	}
}
