package main

import (
	"encoding/json"
	"fmt"
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
