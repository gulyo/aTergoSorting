package main

import (
	"strings"
)

type WeightsArg struct {
	characters []string
	index      int
	literal    *Literal
	result     []float32
}

func weight(arg WeightsArg) WeightsArg {
	if arg.index >= len(arg.characters) {
		return WeightsArg{}
	}
	if arg.literal.Literals == nil {
		return WeightsArg{}
	}
	if literal, isMapContainsKey := arg.literal.Literals[arg.characters[arg.index]]; isMapContainsKey {

		// Checking for multi character literals
		multiCharacterResult := weight(WeightsArg{
			characters: arg.characters,
			index:      arg.index + 1,
			literal:    &literal,
			result:     arg.result,
		})
		if multiCharacterResult.characters != nil {
			return multiCharacterResult
		}

		// Otherwise we use current literal weights
		var res []float32
		res = append(arg.result, literal.Weight...)
		return WeightsArg{
			characters: arg.characters,
			index:      arg.index + 1,
			literal:    arg.literal,
			result:     res,
		}

	}
	return WeightsArg{}
}

func (a Word) CalculateWeights(abc *Literal) WordWeight {
	result := WordWeight{word: a}

	var characters []string = strings.Split(strings.ToUpper(string(a)), "")

	var i int = 0
	var arg WeightsArg = WeightsArg{
		characters: characters,
		index:      i,
		literal:    abc,
		result:     []float32{},
	}
	for i < len(characters) {
		arg = weight(arg)
		arg.literal = abc
		i = arg.index
	}
	result.weights = arg.result

	return result
}
