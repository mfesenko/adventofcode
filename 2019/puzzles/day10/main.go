package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/asteroid"
	"github.com/mfesenko/adventofcode/2019/input"
)

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	asteroidMap := asteroid.NewMap(input)
	detector := asteroid.NewDetector(asteroidMap)
	location, count := detector.FindBestLocation()
	fmt.Printf("Best location: (%v, %v), count: %v\n", location.X, location.Y, count)

	vaporizer := asteroid.NewVaporizer(asteroidMap)
	asteroids := vaporizer.Vaporize(location)
	fmt.Printf("Asteroid #200: (%v, %v)\n", asteroids[199].X, asteroids[199].Y)
}
