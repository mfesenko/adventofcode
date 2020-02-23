package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/image"
	"github.com/mfesenko/adventofcode/2019/input"
)

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	width := 25
	height := 6
	image := image.ParseImage(input[0], width, height)
	fmt.Printf("Part 1: checksum = %v\n", image.CheckSum())
	fmt.Println("Part 2:")
	fmt.Println(image.String())
}
