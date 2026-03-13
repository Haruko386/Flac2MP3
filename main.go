package main

import (
	"Flac2MP3/converter"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.flac")
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")

	input = strings.Trim(input, "\"")
	input = strings.Trim(input, "'")

	if err := converter.Convert(input); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
