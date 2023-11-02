package main

import (
	"bufio"
	"fmt"
	"os"
)

func readNames(location string) chan [ChunkCount]Word {
	channel := make(chan [ChunkCount]Word)

	names, err := os.Open(location)
	if err != nil {
		fmt.Println("Could open names file")
		fmt.Println(err.Error())
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Couldn't close names file")
			fmt.Println(err.Error())
		}
		fmt.Println("Names file closed")
	}(names)

	nameScanner := bufio.NewScanner(names)

	hasText := nameScanner.Scan()
	go func() {
		for hasText {
			var i int = 0
			var chunk [ChunkCount]Word
			for hasText && (i < ChunkCount) {
				chunk[i] = Word(nameScanner.Text())
				i++
				hasText = nameScanner.Scan()
			}
			channel <- chunk
		}
		close(channel)
		fmt.Println("Channel closed")
	}()
	return channel
}
