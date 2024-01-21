package main

import (
	"strings"
)

type WeightsArg struct {
	characters []string
	index      int
	literal    *Literal
	result     []float32
	missing    []string
	multiCheck bool
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
			missing:    arg.missing,
			multiCheck: true,
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
			missing:    arg.missing,
			multiCheck: false,
		}

	} else {
		if arg.multiCheck {
			return WeightsArg{}
		} else {
			return WeightsArg{
				characters: arg.characters,
				index:      arg.index + 1,
				literal:    arg.literal,
				result:     arg.result,
				missing:    append(arg.missing, arg.characters[arg.index]),
				multiCheck: false,
			}
		}
	}
}

func (a Word) CalculateWeights(abc *Literal) WordWeight {
	result := WordWeight{word: a}

	var characters = strings.Split(strings.ToUpper(string(a.text)), "")

	var i = 0
	var arg = WeightsArg{
		characters: characters,
		index:      i,
		literal:    abc,
		result:     []float32{},
		missing:    []string{},
		multiCheck: false,
	}
	for i < len(characters) {
		arg = weight(arg)
		arg.literal = abc
		i = arg.index
	}
	result.weights = arg.result

	return WordWeight{
		word:    a,
		weights: arg.result,
		missing: arg.missing,
	}
}
