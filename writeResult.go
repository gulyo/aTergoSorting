package main

import (
	"log"
	"os"
)

func writeResult(result *[]WordWeight) {
	// create file
	out, err := os.OpenFile(os.Args[3], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	_, err = out.WriteString("Név, Ismeretlen Betűk\n")
	if err != nil {
		log.Fatal(err)
	}

	for _, weighted := range *result {
		_, err := out.WriteString(weighted.String() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
