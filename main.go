package main

import (
	"Flac2MP3/converter"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.flac")
		os.Exit(1)
	}

	input := os.Args[1]

	if err := converter.Convert(input); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
