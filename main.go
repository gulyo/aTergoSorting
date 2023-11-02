package main

import "fmt"

func main() {

	var container [][]WordWeight = sortedParts()

	var result []WordWeight = combineParts(container)

	for _, weighted := range result {
		fmt.Println(weighted)
	}
}
