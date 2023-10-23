package main

import (
	"strings"
)

type Word string

type CalculateWeightsResult struct {
	word    *Word
	weights *[]float32
}

type WeightsArg struct {
	characters *[]string
	index      int
	literal    *Literal
	result     *[]float32
}

func weight(arg *WeightsArg) *WeightsArg {
	if arg.index >= len(*arg.characters) {
		return nil
	}
	if arg.literal.Literals == nil {
		return nil
	}
	if literal, isMapContainsKey := arg.literal.Literals[(*arg.characters)[arg.index]]; isMapContainsKey {

		// Checking for multi character literals
		multiCharacterResult := weight(&WeightsArg{
			characters: arg.characters,
			index:      arg.index + 1,
			literal:    &literal,
			result:     arg.result,
		})
		if multiCharacterResult != nil {
			return multiCharacterResult
		}

		// Otherwise we use current literal weights
		var res []float32
		res = append(*arg.result, literal.Weight...)
		return &WeightsArg{
			characters: arg.characters,
			index:      arg.index + 1,
			literal:    arg.literal,
			result:     &res,
		}

	}
	return nil
}

func (word *Word) CalculateWeights(abc *Literal) *CalculateWeightsResult {
	result := CalculateWeightsResult{word: word}

	var characters []string = strings.Split(strings.ToUpper(string(*word)), "")

	var i int = 0
	var arg *WeightsArg = &WeightsArg{
		characters: &characters,
		index:      i,
		literal:    abc,
		result:     &([]float32{}),
	}
	for i < len(characters) {
		arg = weight(arg)
		arg.literal = abc
		i = arg.index
	}
	result.weights = arg.result

	return &result
}
