package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/arcade"
	"github.com/mfesenko/adventofcode/2019/intcode"
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	game := play(program)
	fmt.Printf("Number of block tiles: %v\n", game.CountTiles(arcade.BlockTile))

	program.Write(0, 2)
	game = play(program)
	fmt.Printf("Score: %v\n", game.Score())
}

func play(program *intcode.Program) *arcade.Game {
	computer := intcode.NewComputer()
	computer.SetProgram(program)
	game := arcade.NewGame(computer)
	game.Play()
	return game
}
