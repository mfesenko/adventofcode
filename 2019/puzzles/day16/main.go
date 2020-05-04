package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/decode"
	"github.com/mfesenko/adventofcode/2019/input"
)

const phaseCount = 100

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	message, err := decode.ParseMessage(input[0])
	if err != nil {
		fmt.Printf("Failed to parse message: %v\n", err)
		return
	}

	result := decode.NewFFTDecoder(phaseCount).Decode(message)
	fmt.Printf("Result of FFT decoder: %v\n", result.String())

	result = decode.NewRealSignalDecoder(phaseCount).Decode(message)
	fmt.Printf("Result of real signal decoder: %v\n", result.String())
}
