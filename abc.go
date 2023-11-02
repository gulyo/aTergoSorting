package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Literal struct {
	Weight   []float32          `json:"weight"`
	Literals map[string]Literal `json:"literals"`
}

func unmarshallAbc(jsonInput []byte) Literal {
	var abc Literal

	err := json.Unmarshal(jsonInput, &abc)
	if err != nil {
		fmt.Println(err)
		return Literal{Weight: []float32{0}}
	}

	return abc
}

func readAbc(location string) Literal {
	abcFile, err := os.ReadFile(location)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	return unmarshallAbc(abcFile)
}
