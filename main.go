package main

import (
	"fmt"
	"os"
)

func main() {

	abcFile, err := os.ReadFile("./resource/abc_hun_extended.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	abc := unmarshallAbc(abcFile)

	var word Word = "Hellyóka"

	result := word.CalculateWeights(&abc)

	fmt.Println(*result.word, " -> ", *result.weights)

}
