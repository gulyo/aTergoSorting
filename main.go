package main

import (
	"fmt"
	"os"
)

func main() {
	var container *[]*[]WordWeight = sortedParts()

	var result *[]WordWeight = combineParts(container)

	if os.Args[3] != "" {
		writeResult(result)
	} else {
		for _, weighted := range *result {
			fmt.Println(weighted)
		}
	}
}
