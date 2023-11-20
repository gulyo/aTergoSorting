package main

import (
	"bufio"
	"fmt"
	"os"
)

func readNames(location string) chan *[]Word {
	channel := make(chan *[]Word)

	names, err := os.Open(location)
	if err != nil {
		fmt.Println("Could open names file")
		panic(err.Error())
	}

	nameScanner := bufio.NewScanner(names)

	hasText := nameScanner.Scan()
	go func() {
		for hasText {
			var i int = 0
			var chunk []Word
			for hasText && (i < ChunkCount) {
				chunk = append(chunk, Word(nameScanner.Text()))
				i++
				hasText = nameScanner.Scan()
			}
			channel <- &chunk
		}
		defer func() { close(channel) }()
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("Couldn't close names file")
				fmt.Println(err.Error())
			}
		}(names)
	}()
	return channel
}
