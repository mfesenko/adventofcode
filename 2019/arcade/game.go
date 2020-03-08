package arcade

import (
	"sync"

	"github.com/mfesenko/adventofcode/2019/async"
	"github.com/mfesenko/adventofcode/2019/math"
)

type (
	// Computer represents a computer that provides instructions to the game
	Computer interface {
		Execute()
		Input() chan int64
		Output() chan int64
	}

	// Game represents an arcade game
	Game struct {
		computer Computer
		tiles    map[math.Point]TileID
		score    int64
		paddleX  int64
	}
)

// NewGame creates a game
func NewGame(computer Computer) *Game {
	return &Game{
		computer: computer,
		tiles:    map[math.Point]TileID{},
	}
}

// CountTiles returns the count of tiles with given id
func (g *Game) CountTiles(tile TileID) int {
	count := 0
	for _, curTile := range g.tiles {
		if curTile == tile {
			count++
		}
	}
	return count
}

// Score returns score
func (g *Game) Score() int64 {
	return g.score
}

// Play starts a game
func (g *Game) Play() {
	done := &sync.WaitGroup{}
	done.Add(1)
	executor := async.NewExecutor(g.computer)
	executor.ExecuteAsync(done)
	for executor.Running() {
		x, y, v, ok := g.readInstructions()
		if !ok {
			break
		}
		g.processInstructions(x, y, v)
	}
	done.Wait()
}

func (g *Game) readInstructions() (int64, int64, int64, bool) {
	instructions := make([]int64, 3)
	for i := range instructions {
		instruction, ok := <-g.computer.Output()
		if !ok {
			return 0, 0, 0, false
		}
		instructions[i] = instruction
	}
	return instructions[0], instructions[1], instructions[2], true
}

func (g *Game) processInstructions(x int64, y int64, v int64) {
	if g.isUpdateScoreInstruction(x, y) {
		g.score = v
		return
	}

	tile := TileID(v)
	g.tiles[math.NewPoint(x, y)] = tile

	if tile == HorizontalPaddleTile {
		g.onPaddleMove(x)
		return
	}

	if tile == BallTile {
		g.onBallMove(x)
		return
	}
}

func (g *Game) isUpdateScoreInstruction(x int64, y int64) bool {
	return x == -1 && y == 0
}

func (g *Game) onPaddleMove(x int64) {
	g.paddleX = x
}

func (g *Game) onBallMove(x int64) {
	if x == g.paddleX {
		g.movePaddle(JoystickNeutral)
	} else if x < g.paddleX {
		g.movePaddle(JoystickLeft)
	} else {
		g.movePaddle(JoystickRight)
	}
}

func (g *Game) movePaddle(position JoystickPosition) {
	g.computer.Input() <- int64(position)
}
