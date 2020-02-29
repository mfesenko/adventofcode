package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/drawing"
	"github.com/mfesenko/adventofcode/2019/intcode"
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	partOne(program)
	partTwo(program)
}

func partOne(program *intcode.Program) {
	robot := createRobot(program)
	robot.ChangeColor(drawing.Black)
	robot.Run()
	fmt.Printf("Covered tiles: %v\n", robot.CoveredTiles())
}

func partTwo(program *intcode.Program) {
	robot := createRobot(program)
	robot.ChangeColor(drawing.White)
	robot.Run()
	fmt.Println(robot.Image().String())
}

func createRobot(program *intcode.Program) *drawing.Robot {
	computer := intcode.NewComputer()
	computer.SetProgram(program)
	return drawing.NewRobot(computer)
}
